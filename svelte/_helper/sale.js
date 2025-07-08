
function countMenuIds(menuIds) {
  if (!menuIds) return {};

  const ids = Array.isArray(menuIds) ? menuIds : [menuIds];
  const counts = {};

  for (const id of ids) {
    counts[id] = (counts[id] || 0) + 1;
  }
  return counts;
}

export function convertMenuChoicesToMenuOptions(rawData) {
     return Object.entries(rawData).map(([id, value]) => {
     const match = value.match(/^(.*)\s+\((\d+)\)$/);
     return {
       id: parseInt(id),
       name: match ? match[1] : value,
       price: match ? parseInt(match[2]) : 0
     };
   });
}
 
export function convertSalesToTodaySales(sales, menuOptions) {
   const result = [];
 
   sales.forEach((sale) => {
     const qtyMap = countMenuIds(sale.menuIds);
 
     Object.entries(qtyMap).forEach(([menuId, qty]) => {
       const menu = menuOptions.find((m) => m.id === parseInt(menuId));
       if (menu) {
         const totalPrice = menu.price * qty;
         const time = new Date(sale.createdAt * 1000).toLocaleTimeString("id-ID", {
           hour: "2-digit",
           minute: "2-digit",
           hour12: false
         });
 
         result.push({
           id: parseInt(sale.id),
           item: menu.name,
           qty,
           price: totalPrice,
           time
         });
       }
     });
   });
 
   return result;
 }