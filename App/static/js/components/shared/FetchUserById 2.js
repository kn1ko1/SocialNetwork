export const fetchUserById = async userId => {
  try {
    const response = await fetch(`http://localhost:8080/api/users/${userId}`);
    const data = await response.json();
    return data;
  } catch (error) {
    console.error(`Error fetching username for user ${userId}:`, error);
    return null;
  }
};