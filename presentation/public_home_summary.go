package presentation

import (
	"sort"
	"time"

	"kostjc/model/mProperty/rqProperty"
)

type PublicUnpaidTenant struct {
	TenantName string
	RoomName   string
	LastDate   string
	DaysAgo    int
}

type PublicRoomNotice struct {
	RoomName   string
	LastTenant string
	LastDate   string
	DaysAgo    int
}

type PublicLoginSummary struct {
	UnpaidTenants []PublicUnpaidTenant
	Rooms         []PublicRoomNotice
	Occupants     []PublicOccupant
}

type PublicOccupant struct {
	TenantName string
	RoomNames  []string
}

func buildPublicLoginSummary(unpaid []rqProperty.UnpaidBookingTenant, rooms []rqProperty.AvailableRoom, occupants []rqProperty.CurrentOccupant) PublicLoginSummary {
	return buildPublicLoginSummaryAt(unpaid, rooms, occupants, time.Now().UTC())
}

func buildPublicLoginSummaryAt(unpaid []rqProperty.UnpaidBookingTenant, rooms []rqProperty.AvailableRoom, occupants []rqProperty.CurrentOccupant, now time.Time) PublicLoginSummary {
	seenUnpaid := map[string]struct{}{}
	unpaidOut := make([]PublicUnpaidTenant, 0, len(unpaid))
	usedRooms := map[string]struct{}{}
	today := now.UTC().Format(time.DateOnly)
	for _, item := range unpaid {
		if item.TenantName == "" || item.RoomName == "" {
			continue
		}
		if dateLess(today, item.DateStart) {
			rooms = append(rooms, rqProperty.AvailableRoom{
				RoomName:    item.RoomName,
				LastTenant:  item.TenantName,
				AvailableAt: item.DateStart,
			})
			continue
		}
		key := item.TenantName + "\x00" + item.RoomName
		if _, ok := seenUnpaid[key]; ok {
			continue
		}
		seenUnpaid[key] = struct{}{}
		lastDate := item.DateEnd
		usedRooms[item.RoomName] = struct{}{}
		unpaidOut = append(unpaidOut, PublicUnpaidTenant{
			TenantName: item.TenantName,
			RoomName:   item.RoomName,
			LastDate:   lastDate,
			DaysAgo:    unpaidDaysAt(item.DateStart, lastDate, now),
		})
	}

	seenRooms := map[string]struct{}{}
	roomsOut := make([]PublicRoomNotice, 0, len(rooms))
	for _, item := range rooms {
		if item.RoomName == "" || item.LastTenant == "" {
			continue
		}
		key := item.RoomName + "\x00" + item.LastTenant + "\x00" + item.AvailableAt
		if _, ok := seenRooms[key]; ok {
			continue
		}
		seenRooms[key] = struct{}{}
		lastDate := item.AvailableAt
		usedRooms[item.RoomName] = struct{}{}
		roomsOut = append(roomsOut, PublicRoomNotice{
			RoomName:   item.RoomName,
			LastTenant: item.LastTenant,
			LastDate:   lastDate,
			DaysAgo:    signedDaysAt(lastDate, now),
		})
	}

	seenOccupants := map[string]map[string]struct{}{}
	occupantsOut := make([]PublicOccupant, 0, len(occupants))
	for _, item := range occupants {
		if item.TenantName == "" || item.RoomName == "" {
			continue
		}
		if _, ok := usedRooms[item.RoomName]; ok {
			continue
		}
		if seenOccupants[item.TenantName] == nil {
			seenOccupants[item.TenantName] = map[string]struct{}{}
		}
		if _, ok := seenOccupants[item.TenantName][item.RoomName]; ok {
			continue
		}
		seenOccupants[item.TenantName][item.RoomName] = struct{}{}
		found := false
		for i := range occupantsOut {
			if occupantsOut[i].TenantName == item.TenantName {
				occupantsOut[i].RoomNames = append(occupantsOut[i].RoomNames, item.RoomName)
				found = true
				break
			}
		}
		if !found {
			occupantsOut = append(occupantsOut, PublicOccupant{
				TenantName: item.TenantName,
				RoomNames:  []string{item.RoomName},
			})
		}
	}

	for i := range occupantsOut {
		sort.Strings(occupantsOut[i].RoomNames)
	}
	sort.Slice(occupantsOut, func(i, j int) bool {
		if len(occupantsOut[i].RoomNames) == 0 {
			return len(occupantsOut[j].RoomNames) > 0
		}
		if len(occupantsOut[j].RoomNames) == 0 {
			return true
		}
		if occupantsOut[i].RoomNames[0] != occupantsOut[j].RoomNames[0] {
			return occupantsOut[i].RoomNames[0] < occupantsOut[j].RoomNames[0]
		}
		return occupantsOut[i].TenantName < occupantsOut[j].TenantName
	})

	sort.Slice(unpaidOut, func(i, j int) bool {
		return dateLess(unpaidOut[i].LastDate, unpaidOut[j].LastDate)
	})
	sort.Slice(roomsOut, func(i, j int) bool {
		return dateLess(roomsOut[i].LastDate, roomsOut[j].LastDate)
	})

	return PublicLoginSummary{
		UnpaidTenants: unpaidOut,
		Rooms:         roomsOut,
		Occupants:     occupantsOut,
	}
}

func dateLess(left, right string) bool {
	l, errL := time.Parse(time.DateOnly, left)
	r, errR := time.Parse(time.DateOnly, right)
	if errL != nil && errR != nil {
		return left < right
	}
	if errL != nil {
		return false
	}
	if errR != nil {
		return true
	}
	return l.Before(r)
}

func lateDays(dateStr string) int {
	if dateStr == "" {
		return 0
	}
	dt, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return 0
	}
	now := time.Now()
	dt = time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, time.UTC)
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	if now.Before(dt) {
		return 0
	}
	return int(now.Sub(dt).Hours() / 24)
}

func signedDays(dateStr string) int {
	return signedDaysAt(dateStr, time.Now().UTC())
}

func signedDaysAt(dateStr string, now time.Time) int {
	if dateStr == "" {
		return 0
	}
	dt, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return 0
	}
	now = now.UTC()
	dt = time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, time.UTC)
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return int(now.Sub(dt).Hours() / 24)
}

func unpaidDays(startDate, endDate string) int {
	return unpaidDaysAt(startDate, endDate, time.Now().UTC())
}

func unpaidDaysAt(startDate, endDate string, now time.Time) int {
	if startDate != "" {
		return signedDaysAt(startDate, now)
	}
	return signedDaysAt(endDate, now)
}
