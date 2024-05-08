export const fetchUsername = async userId => {
  try {
    const response = await fetch(`http://localhost:8080/api/users/${userId}`);
    const data = await response.json();
    return data.username;
  } catch (error) {
    console.error("Error fetching username:", error);
    return null;
  }
};