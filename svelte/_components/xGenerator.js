
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