
function countMenuIds(menuIds) {
  if (!menuIds) return {};

  const ids = Array.isArray(menuIds) ? menuIds : [menuIds];
  const counts = {};

  for (const id of ids) {
    counts[id] = (counts[id] || 0) + 1;
  }
  return counts;
}

export function parseMenuMapToOptions(mapData) {
  const menuOptions = [];

  Object.keys(mapData).forEach(key => {
    const value = mapData[key];
    
    const match = value.match(/^(.*?)\s\((\d+)\)\s\((.*?)\)$/);

    if (match) {
      const name = match[1].trim();
      const price = parseInt(match[2], 10);
      const imageUrl = match[3].trim();

      menuOptions.push({
        id: parseInt(key, 10),
        name,
        price,
        imageUrl,
      });
    }
  });

  return menuOptions;
}
 
export function convertSalesToTodaySales(sales, menuOptions) {
   const todaySales = [];
 
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
 
         todaySales.push({
           id: parseInt(sale.id),
           item: menu.name,
           qty,
           price: totalPrice,
           time
         });
       }
     });
   });
 
   return todaySales;
 }

 export function formatCurrency(amount) {
  return new Intl.NumberFormat('id-ID').format(amount);
}

export function getBuyerOptions(sales, tenants) {
  const result = {};
  
  const today = new Date().toISOString().split('T')[0];
  
  for (const sale of sales) {
    const saleId = sale[0];
    const buyerName = sale[3];
    const tenantId = sale[2];
    const totalPrice = sale[14];
    const salesDate = sale[15];
    const paymentStatus = sale[6];
    
    if (salesDate !== today) {
      continue;
    }
    
    if (paymentStatus !== 'Unpaid') {
      continue;
    }

    let label = `${saleId} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;

    if (tenantId > 0) {
      const fullTenantInfo = tenants[tenantId];
      if (fullTenantInfo) {
        const [tenantName, teleNumber, waNumber] = fullTenantInfo.split(' / ');
        const info = waNumber ? `${tenantName} / ${waNumber}` : tenantName;
        label = `${saleId} / ${info} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;
      } else {
        label = `${saleId} / ${tenantId} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;
      }
    }

    result[saleId] = label;
  }
  
  return result;
}

export function getBuyerOptionsForUnpaidItems(sales, tenants) {
  const result = {};
  
  
  for (const sale of sales) {
    const saleId = sale[0];
    const buyerName = sale[3];
    const tenantId = sale[2];
    const totalPrice = sale[14];
    const paymentStatus = sale[6];
    
    
    if (paymentStatus !== 'Unpaid') {
      continue;
    }

    let label = `${saleId} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;

    if (tenantId > 0) {
      const fullTenantInfo = tenants[tenantId];
      if (fullTenantInfo) {
        const [tenantName, teleNumber, waNumber] = fullTenantInfo.split(' / ');
        const info = waNumber ? `${tenantName} / ${waNumber}` : tenantName;
        label = `${saleId} / ${info} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;
      } else {
        label = `${saleId} / ${tenantId} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;
      }
    }

    result[saleId] = label;
  }
  
  return result;
}


export function parseSalesToTodaySales(sales, menusAsObject) {
  const salesObj = {};
  
  const today = new Date().toISOString().split('T')[0];

  for (const sale of sales) {
    const salesDate = sale[15];
    
    if (salesDate !== today) continue;

    const menuIds = sale[4];
    const totalPrice = sale[14];
    const perItemPrice = Math.floor(totalPrice / menuIds.length);
    const timeString = new Date(sale[17] * 1000).toLocaleTimeString("id-ID", {
      hour: "2-digit",
      minute: "2-digit", 
      hour12: false
    });

    for (const id of menuIds) {
      if (salesObj[id]) {
        salesObj[id].qty += 1;
        salesObj[id].total += menusAsObject[id]?.price || perItemPrice;
        salesObj[id].times = timeString;
      } else {
        salesObj[id] = {
          id: id,
          name: menusAsObject[id]?.name || `Menu ${id}`,
          qty: 1,
          total: menusAsObject[id]?.price || perItemPrice,
          times: timeString
        };
      }
    }
  }
  
  return Object.values(salesObj);
}



export function convertMenusToObject(menus) {
  const obj = {};
  for (const menu of menus) {
    obj[menu.id] = menu;
  }
  return obj;
}

export function parseSalesToTodayPayment(salesData, tenants) {
  const todayPayments = [];
  
  if (!Array.isArray(salesData) || salesData.length === 0) {
      console.log('Data sales kosong atau tidak valid');
      return todayPayments;
  }
  
  salesData.forEach((salesArray, index) => {
      try {
          if (!Array.isArray(salesArray) || salesArray.length < 21) {
              console.warn(`Sales data index ${index} tidak valid atau kurang lengkap`);
              return;
          }
          
          const status = salesArray[6];
          
          if (status === "Paid" || status === "Overpaid") {
              const payment = {
                  id: parseInt(salesArray[0]) || 0,
                  customer: getCleanCustomerName(salesArray, tenants) || '',
                  amount: calculateSaleTotal(salesArray)|| 0,
                  method: salesArray[5] || '',
                  status: salesArray[6] || '',
                  time: formatTimestamp(salesArray[18])
              };
              
              todayPayments.push(payment);
              console.log(`Added payment: ID ${payment.id}, Customer: ${payment.customer}, Amount: ${payment.amount}`);
          }
      } catch (error) {
          console.error(`Error parsing sales data at index ${index}:`, error);
      }
  });
  
  return todayPayments;
}

export function parseSalesToOverpaid(salesData, tenants) {
  const overpaidData = [];
  
  if (!Array.isArray(salesData) || salesData.length === 0) {
      console.log('Data sales kosong atau tidak valid');
      return overpaidData;
  }
  
  salesData.forEach((salesArray, index) => {
      try {
          if (!Array.isArray(salesArray) || salesArray.length < 21) {
              console.warn(`Sales data index ${index} tidak valid atau kurang lengkap`);
              return;
          }
          
          const status = salesArray[6];
          
          if (status === "Overpaid") {
              const paidAmount = calculateSaleTotal(salesArray) || 0;
              const totalAmount = parseInt(salesArray[14]) || 0;
              const excess = Math.max(0, paidAmount - totalAmount);
              
              const overpaid = {
                  id: parseInt(salesArray[0]) || 0,
                  customer: getCleanCustomerName(salesArray, tenants) || '',
                  excess: excess,
                  method: salesArray[5] || '',
                  time: formatTimestamp(salesArray[18]),
                  date: salesArray[16] || ''
              };
              
              overpaidData.push(overpaid);
              console.log(`Added overpaid: ID ${overpaid.id}, Customer: ${overpaid.customer}, Excess: ${overpaid.excess}`);
          }
      } catch (error) {
          console.error(`Error parsing overpaid data at index ${index}:`, error);
      }
  });
  
  return overpaidData;
}


export function parseSalesToUnpaid(salesData, menus, tenants) {
  const unpaidData = [];
  
  if (!Array.isArray(salesData) || salesData.length === 0) {
      console.log('Data sales kosong atau tidak valid');
      return unpaidData;
  }
  
  salesData.forEach((salesArray, index) => {
      try {
          if (!Array.isArray(salesArray) || salesArray.length < 21) {
              console.warn(`Sales data index ${index} tidak valid atau kurang lengkap`);
              return;
          }
          
          const status = salesArray[6];
          
          if (status === "Unpaid") {
              const unpaid = {
                  id: parseInt(salesArray[0]) || 0,
                  customer: getCleanCustomerName(salesArray, tenants) || '',
                  item: getMenuStringFromSale(salesArray, menus) || '',
                  amount: parseInt(salesArray[14]) || 0,
                  date: salesArray[15] || ''
              };
              
              unpaidData.push(unpaid);
              console.log(`Added unpaid: ID ${unpaid.id}, Customer: ${unpaid.customer}, Amount: ${unpaid.amount}`);
          }
      } catch (error) {
          console.error(`Error parsing unpaid data at index ${index}:`, error);
      }
  });
  
  return unpaidData;
}

function formatTimestamp(timestamp) {
  try {
      if (!timestamp || typeof timestamp !== 'number') {
          return new Date().toLocaleTimeString('id-ID');
      }
      
      const date = new Date(timestamp > 1000000000000 ? timestamp : timestamp * 1000);
      return date.toLocaleTimeString('id-ID', {
          hour: '2-digit',
          minute: '2-digit'
      });
  } catch (error) {
      console.error('Error formatting timestamp:', error);
      return new Date().toLocaleTimeString('id-ID');
  }
}

function calculateSaleTotal(saleArray) {
  try {
      if (!Array.isArray(saleArray) || saleArray.length < 20) {
          return 0;
      }
      const transferAmount = parseInt(saleArray[7]) || 0;
      const qrisAmount = parseInt(saleArray[8]) || 0;
      const cashAmount = parseInt(saleArray[9]) || 0;
      const debtAmount = parseInt(saleArray[10]) || 0;
      const topupAmount = parseInt(saleArray[11]) || 0;
      const donationAmount = parseInt(saleArray[12]) || 0;

      return transferAmount + qrisAmount + cashAmount + debtAmount + topupAmount + donationAmount;
      
  } catch (error) {
      console.error('Error calculating sale total:', error);
      return 0;
  }
}


function getCustomerName(saleArray, tenants) {
  try {
      if (!Array.isArray(saleArray) || saleArray.length < 21) {
          return '';
      }

      const tenantId = parseInt(saleArray[2]) || 0;
      const buyerName = saleArray[3] || '';

      if (tenantId === 0) {
          return buyerName;
      }

      if (tenants && tenants[tenantId]) {
          return tenants[tenantId];
      }
      return `Tenant ${tenantId} (Not Found)`;

  } catch (error) {
      console.error('Error getting customer name:', error);
      return '';
  }
}

function getCleanCustomerName(saleArray, tenants) {
  try {
      const fullName = getCustomerName(saleArray, tenants);
      
      const cleanName = fullName.split(' / ')[0].split(' @')[0].trim();
      
      return cleanName;

  } catch (error) {
      console.error('Error getting clean customer name:', error);
      return '';
  }
}



function getMenuString(menuIds, menus) {
  try {
      if (!Array.isArray(menuIds) || menuIds.length === 0) {
          return '';
      }

      if (!menus) {
          return '';
      }

      let menusObject = menus;
      
      if (Array.isArray(menus)) {
          menusObject = convertMenusToObject(menus);
      }

      const menuCount = {};
      menuIds.forEach(menuId => {
          const id = parseInt(menuId);
          menuCount[id] = (menuCount[id] || 0) + 1;
      });

      const menuStrings = [];
      
      for (const menuId of Object.keys(menuCount)) {
          const id = parseInt(menuId);
          const quantity = menuCount[id];
          
          let menuName = `Menu ${id}`;
          if (menusObject[id]?.name) {
              menuName = menusObject[id].name;
          }

          menuStrings.push(quantity > 1 ? `${menuName} x${quantity}` : menuName);
      }

      return menuStrings.join(', ');

  } catch (error) {
      console.error('Error getting menu string:', error);
      return '';
  }
}

function getMenuStringFromSale(saleArray, menus) {
  try {
      if (!Array.isArray(saleArray) || saleArray.length < 21) {
          return '';
      }
      const menuIds = saleArray[4];
      
      return getMenuString(menuIds, menus);

  } catch (error) {
      console.error('Error getting menu string from sale:', error);
      return '';
  }
}
