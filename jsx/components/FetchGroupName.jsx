export const fetchGroupName = async (objectId) => {
    try {
        const response = await fetch(`http://localhost:8080/api/groups/${objectId}`);
        const data = await response.json();
        return data.title;
    } catch (error) {
        console.error("Error fetching group name:", error);
        return null;
    }
};