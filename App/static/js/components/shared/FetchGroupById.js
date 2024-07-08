export const fetchGroupById = async objectId => {
  try {
    const response = await fetch(`http://localhost:8080/api/groups/${objectId}`);
    const data = await response.json();
    return data;
  } catch (error) {
    console.error(`Error fetching group ${objectId}:`, error);
    return null;
  }
};