
/**
 * Generates a random string of the given length from a set of alphanumeric characters.
 * @param {number} length The length of the string to generate.
 * @returns {string} The generated random string.
 */
export function RandString(length) {
  const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  let result = '';
  
  for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * characters.length);
    result += characters[randomIndex];
  }
  
  return result;
}

/**
   * @description Returns the relative day label based on the given date string (YYYY-MM-DD)
   * @param {string} dateStr - The date string in format YYYY-MM-DD
   * @returns {string} - The relative label like 'H', 'H-1', 'H+3', etc.
   */
export function GetRelativeDayLabel(dateStr) {
  const inputDate = new Date(dateStr);
  const currentDate = new Date();
  
  inputDate.setHours(0, 0, 0, 0);
  currentDate.setHours(0, 0, 0, 0);

  const msInDay = 1000 * 60 * 60 * 24; // @ts-ignore
  const diffDays = Math.round((currentDate - inputDate) / msInDay);

  if (diffDays === 0) return 'H';
  if (diffDays > 0) return `H+${diffDays}`;
  
  return `H${diffDays}`;
}