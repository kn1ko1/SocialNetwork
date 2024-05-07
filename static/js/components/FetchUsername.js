export const fetchUsername = async senderId => {
  try {
    const response = await fetch(`http://localhost:8080/api/users/${senderId}`);
    const data = await response.json();
    return data.username;
  } catch (error) {
    console.error("Error fetching username:", error);
    return null;
  }
};