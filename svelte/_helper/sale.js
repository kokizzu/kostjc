
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


// export function getBuyerOptions(sales, tenants) {
//   const result = {};
//   for (const sale of sales) {
//     const saleId = sale[0];
//     const buyerName = sale[3];
//     const tenantId = sale[2];
//     const totalPrice = sale[13];

//     let label = `#${saleId} / ${buyerName} / (Rp ${formatCurrency(totalPrice)})`;

//     if (tenantId > 0) {
//       const tenantName = tenants[tenantId];
//       if (tenantName) {
//         label = `#${saleId} / ${tenantName} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;
//       } else {
//         label = `#${saleId} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;
//       }
//     }

//     result[saleId] = label;
//   }
//   return result;
// }

// export function getBuyerOptions(sales, tenants) {
//   const result = {};
//   for (const sale of sales) {
//     const saleId = sale[0];
//     const buyerName = sale[3];
//     const tenantId = sale[2];
//     const totalPrice = sale[13];

//     let label = ` #${saleId} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;

//     if (tenantId > 0) {
//       const fullTenantInfo = tenants[tenantId]; // contoh: "Ardianyur Arkom / +62 812-3456-7890"
//       if (fullTenantInfo) {
//         // ambil hanya nama + nomor WA, potong berdasarkan ' / '
//         const [tenantName, teleNumber, waNumber] = fullTenantInfo.split(' / ');
//         const info = waNumber ? `${tenantName} / ${waNumber}` : tenantName;
//         label = `#${saleId} / ${info} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;
//       } else {
//         label = `#${saleId} / ${tenantId} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;
//       }
//     }

//     result[saleId] = label;
//   }
//   return result;
// }

export function getBuyerOptions(sales, tenants) {
  const result = {};
  
  // Get today's date in YYYY-MM-DD format
  const today = new Date().toISOString().split('T')[0];
  
  for (const sale of sales) {
    const saleId = sale[0];
    const buyerName = sale[3];
    const tenantId = sale[2];
    const totalPrice = sale[13]; // totalPriceIDR is at index 13
    const salesDate = sale[14]; // salesDate is at index 14
    const paymentStatus = sale[6]; // paymentStatus is at index 6
    
    // Filter: hanya tampilkan data hari ini dan payment status Unpaid
    if (salesDate !== today) {
      continue; // Skip jika bukan hari ini
    }
    
    if (paymentStatus !== 'Unpaid') {
      continue; // Skip jika bukan status Unpaid
    }

    let label = ` #${saleId} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;

    if (tenantId > 0) {
      const fullTenantInfo = tenants[tenantId]; // contoh: "Ardianyur Arkom / +62 812-3456-7890"
      if (fullTenantInfo) {
        // ambil hanya nama + nomor WA, potong berdasarkan ' / '
        const [tenantName, teleNumber, waNumber] = fullTenantInfo.split(' / ');
        const info = waNumber ? `${tenantName} / ${waNumber}` : tenantName;
        label = `#${saleId} / ${info} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;
      } else {
        label = `#${saleId} / ${tenantId} / ${buyerName} (Rp ${formatCurrency(totalPrice)})`;
      }
    }

    result[saleId] = label;
  }
  
  return result;
}


// export function parseSalesPerItem(sales, menusAsObject) {
//   const todaySales = [];
  
//   // Get today's date in YYYY-MM-DD format
//   const today = new Date().toISOString().split('T')[0];

//   for (const sale of sales) {
//     const salesDate = sale[14]; // salesDate at index 10
    
//     // Filter: hanya proses pembelian hari ini
//     if (salesDate !== today) {
//       continue; // Skip jika bukan hari ini
//     }

//     const menuIds = sale[4];
//     const totalPrice = sale[13]; // totalPriceIDR at index 13

//     const perItemPrice = Math.floor(totalPrice / menuIds.length);

//     for (const id of menuIds) {
//       if (!todaySales[id]) {
//         todaySales[id] = {
//           name: menusAsObject[id]?.name,
//           qty: 0,
//           total: 0,
//           times: '',
//         };
//       }

//       todaySales[id].qty += 1;
//       todaySales[id].total += menusAsObject[id]?.price || perItemPrice;
//       todaySales[id].times = new Date(sale[17] * 1000).toLocaleTimeString("id-ID", {
//           hour: "2-digit",
//           minute: "2-digit",
//           hour12: false
//         });
//     }
//   }
//   return todaySales;
// }

export function parseSalesToTodaySales(sales, menusAsObject) {
  const salesObj = {};
  
  const today = new Date().toISOString().split('T')[0];

  for (const sale of sales) {
    const salesDate = sale[14];
    
    if (salesDate !== today) continue;

    const menuIds = sale[4];
    const totalPrice = sale[13];
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
  
  // Validasi input
  if (!Array.isArray(salesData) || salesData.length === 0) {
      console.log('Data sales kosong atau tidak valid');
      return todayPayments;
  }
  
  salesData.forEach((salesArray, index) => {
      try {
          // Pastikan salesArray adalah array dan memiliki minimal 20 element
          if (!Array.isArray(salesArray) || salesArray.length < 20) {
              console.warn(`Sales data index ${index} tidak valid atau kurang lengkap`);
              return;
          }
          
          // Ambil status dari index 6
          const status = salesArray[6];
          
          // Filter hanya yang berstatus "Paid"
          if (status === "Paid" || status === "Overpaid") {
              // Parse data ke format Payment
              const payment = {
                  id: parseInt(salesArray[0]) || 0, // Index 0: ID
                  customer: getCleanCustomerName(salesArray, tenants) || '', // Index 1: Customer name
                  amount: calculateSaleTotal(salesArray)|| 0, // Index 13: Total amount
                  method: salesArray[5] || '', // Index 5: Payment method
                  status: salesArray[6] || '', // Index 6: Status
                  time: formatTimestamp(salesArray[17]) // Index 17: Timestamp
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
  
  // Validasi input
  if (!Array.isArray(salesData) || salesData.length === 0) {
      console.log('Data sales kosong atau tidak valid');
      return overpaidData;
  }
  
  salesData.forEach((salesArray, index) => {
      try {
          // Pastikan salesArray adalah array dan memiliki minimal 20 element
          if (!Array.isArray(salesArray) || salesArray.length < 20) {
              console.warn(`Sales data index ${index} tidak valid atau kurang lengkap`);
              return;
          }
          
          // Ambil status dari index 6
          const status = salesArray[6];
          
          // Filter hanya yang berstatus "Overpaid"
          if (status === "Overpaid") {
              // Untuk overpaid, kita perlu hitung excess amount
              // Asumsi: index 8 adalah amount yang dibayar, index 13 adalah total yang seharusnya
              const paidAmount = calculateSaleTotal(salesArray) || 0;
              const totalAmount = parseInt(salesArray[13]) || 0;
              const excess = Math.max(0, paidAmount - totalAmount);
              
              // Parse data ke format Overpaid
              const overpaid = {
                  id: parseInt(salesArray[0]) || 0, // Index 0: ID
                  customer: getCleanCustomerName(salesArray, tenants) || '', // Index 1: Customer name
                  excess: excess, // Kelebihan pembayaran
                  date: salesArray[15] || '' // Index 15: Date
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
  
  // Validasi input
  if (!Array.isArray(salesData) || salesData.length === 0) {
      console.log('Data sales kosong atau tidak valid');
      return unpaidData;
  }
  
  salesData.forEach((salesArray, index) => {
      try {
          // Pastikan salesArray adalah array dan memiliki minimal 20 element
          if (!Array.isArray(salesArray) || salesArray.length < 20) {
              console.warn(`Sales data index ${index} tidak valid atau kurang lengkap`);
              return;
          }
          
          // Ambil status dari index 6
          const status = salesArray[6];
          
          // Filter hanya yang berstatus "Unpaid"
          if (status === "Unpaid") {
              // Parse data ke format Unpaid
              const unpaid = {
                  id: parseInt(salesArray[0]) || 0, // Index 0: ID
                  customer: getCleanCustomerName(salesArray, tenants) || '', // Index 1: Customer name
                  item: getMenuStringFromSale(salesArray, menus) || '', // Index 3: Item description (pembeli pertama, etc)
                  amount: parseInt(salesArray[13]) || 0, // Index 13: Total amount
                  date: salesArray[14] || '' // Index 14: Date
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
      // Validasi array
      if (!Array.isArray(saleArray) || saleArray.length < 20) {
          return 0;
      }

      // Ambil nilai dari setiap metode pembayaran dan jumlahkan
      const transferAmount = parseInt(saleArray[7]) || 0;  // TransferIDR
      const qrisAmount = parseInt(saleArray[8]) || 0;      // QrisIDR
      const cashAmount = parseInt(saleArray[9]) || 0;      // CashIDR
      const debtAmount = parseInt(saleArray[10]) || 0;     // DebtIDR
      const topupAmount = parseInt(saleArray[11]) || 0;    // TopupIDR
      const donationAmount = parseInt(saleArray[12]) || 0; // Donation

      return transferAmount + qrisAmount + cashAmount + debtAmount + topupAmount + donationAmount;
      
  } catch (error) {
      console.error('Error calculating sale total:', error);
      return 0;
  }
}


function getCustomerName(saleArray, tenants) {
  try {
      // Validasi array
      if (!Array.isArray(saleArray) || saleArray.length < 20) {
          return '';
      }

      const tenantId = parseInt(saleArray[2]) || 0; // Index 2: TenantId
      const buyerName = saleArray[3] || ''; // Index 3: BuyerName

      // Jika tenantId = 0, return buyerName
      if (tenantId === 0) {
          return buyerName;
      }

      // Jika tenantId != 0, return nama dari tenants object
      if (tenants && tenants[tenantId]) {
          return tenants[tenantId];
      }

      // Fallback jika tenant tidak ditemukan
      return `Tenant ${tenantId} (Not Found)`;

  } catch (error) {
      console.error('Error getting customer name:', error);
      return '';
  }
}

function getCleanCustomerName(saleArray, tenants) {
  try {
      const fullName = getCustomerName(saleArray, tenants);
      
      // Potong nama jika ada info tambahan (sebelum " / " atau " @")
      const cleanName = fullName.split(' / ')[0].split(' @')[0].trim();
      
      return cleanName;

  } catch (error) {
      console.error('Error getting clean customer name:', error);
      return '';
  }
}



function getMenuString(menuIds, menus) {
  try {
      // Validasi input
      if (!Array.isArray(menuIds) || menuIds.length === 0) {
          return '';
      }

      if (!menus) {
          return '';
      }

      // Untuk performance terbaik, gunakan Object of Object
      let menusObject = menus;
      
      // Convert array ke object hanya jika diperlukan (performance cost)
      if (Array.isArray(menus)) {
          menusObject = convertMenusToObject(menus);
      }

      // Hitung quantity untuk setiap menu - O(n) where n = menuIds.length
      const menuCount = {};
      menuIds.forEach(menuId => {
          const id = parseInt(menuId);
          menuCount[id] = (menuCount[id] || 0) + 1;
      });

      // Buat string menu dengan quantity - O(k) where k = unique menus
      const menuStrings = [];
      
      for (const menuId of Object.keys(menuCount)) {
          const id = parseInt(menuId);
          const quantity = menuCount[id];
          
          // O(1) lookup dengan object, O(n) dengan array
          let menuName = `Menu ${id}`;
          if (menusObject[id]?.name) {
              menuName = menusObject[id].name;
          }

          // Format berdasarkan quantity
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
      // Validasi array
      if (!Array.isArray(saleArray) || saleArray.length < 20) {
          return '';
      }

      // Ambil menuIds dari index 4
      const menuIds = saleArray[4];
      
      return getMenuString(menuIds, menus);

  } catch (error) {
      console.error('Error getting menu string from sale:', error);
      return '';
  }
}
