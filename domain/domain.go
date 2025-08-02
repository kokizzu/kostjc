package domain

import (
	"net"
	"time"

	chBuffer "github.com/kokizzu/ch-timed-buffer"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/M"
	"github.com/rs/zerolog"

	"kostjc/conf"
	"kostjc/model/mAuth"
	"kostjc/model/mAuth/saAuth"
	"kostjc/model/mCafe"
	"kostjc/model/mCafe/saCafe"
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/saProperty"
	"kostjc/model/xMailer"
)

type Domain struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter

	PropOltp *Tt.Adapter
	PropOlap *Ch.Adapter

	Mailer xMailer.Mailer

	IsBgSvc bool // long-running program

	// timed buffer
	authLogs       *chBuffer.TimedBuffer
	userLogs       *chBuffer.TimedBuffer
	tenantLogs     *chBuffer.TimedBuffer
	locationLogs   *chBuffer.TimedBuffer
	facilityLogs   *chBuffer.TimedBuffer
	buildingLogs   *chBuffer.TimedBuffer
	roomLogs       *chBuffer.TimedBuffer
	bookingLogs    *chBuffer.TimedBuffer
	paymentLogs    *chBuffer.TimedBuffer
	stockLogs      *chBuffer.TimedBuffer
	menuLogs       *chBuffer.TimedBuffer
	saleLogs       *chBuffer.TimedBuffer
	wifiDeviceLogs *chBuffer.TimedBuffer

	// logger
	Log *zerolog.Logger

	// list of superadmin emails
	Superadmins M.SB

	// web related, need to be here so Mailer/CLI also know the domain/urls
	WebCfg conf.WebConf
}

// will run in background if background service
func (d *Domain) runSubtask(subTask func()) {
	if d.IsBgSvc {
		go subTask()
	} else {
		subTask()
	}
}

func (d *Domain) InitTimedBuffer() {
	d.authLogs = chBuffer.NewTimedBuffer(d.AuthOlap.DB, 100_000, 1*time.Second, saAuth.Preparators[mAuth.TableActionLogs])
	d.userLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saAuth.Preparators[mAuth.TableUserLogs])
	d.tenantLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saAuth.Preparators[mAuth.TableTenantLogs])
	d.locationLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TableLocationLogs])
	d.facilityLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TableFacilityLogs])
	d.buildingLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TableBuildingLogs])
	d.roomLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TableRoomLogs])
	d.bookingLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TableBookingLogs])
	d.paymentLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TablePaymentLogs])
	d.stockLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TableStockLogs])
	d.menuLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saCafe.Preparators[mCafe.TableMenuLogs])
	d.saleLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saCafe.Preparators[mCafe.TableSaleLogs])
	d.wifiDeviceLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TableWifiDeviceLogs])
}

func (d *Domain) WaitTimedBufferFinalFlush() {
	<-d.authLogs.WaitFinalFlush
	d.Log.Debug().Msg(`timed buffer flushed`)
}

var defaultIP4 = net.ParseIP(`0.0.0.0`).To4()
var defaultIP6 = net.ParseIP(`0:0:0:0:0:0:0:0`).To16()

func (d *Domain) InsertActionLog(in *RequestCommon, out *ResponseCommon) bool {
	ip := net.ParseIP(in.IpAddress)
	ip4 := ip.To4()
	if ip4 == nil {
		ip4 = defaultIP4
	}
	ip6 := ip.To16()
	if ip6 == nil {
		ip6 = defaultIP6
	}
	row := saAuth.ActionLogs{
		CreatedAt:  in.TimeNow(),
		RequestId:  in.RequestId,
		ActorId:    out.actor,
		Action:     in.Action,
		StatusCode: int16(out.StatusCode),
		Traces:     out.Traces(),
		Error:      out.Error,
		IpAddr4:    ip4,
		IpAddr6:    ip6,
		UserAgent:  in.UserAgent,
		Lat:        in.Lat,
		Long:       in.Long,
		Latency:    in.Latency(),
		RefId:      in.RefId,
	}
	return d.authLogs.Insert([]any{
		row.CreatedAt,
		row.RequestId,
		row.ActorId,
		row.Action,
		row.StatusCode,
		row.Traces,
		row.Error,
		row.IpAddr4,
		row.IpAddr6,
		row.UserAgent,
		row.Lat,
		row.Long,
		row.Latency,
		row.RefId,
	})
}

func (d *Domain) CloseTimedBuffer() {
	go d.authLogs.Close()
	d.WaitTimedBufferFinalFlush()
}
