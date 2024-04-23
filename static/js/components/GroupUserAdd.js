export async function AddGroupUser({
  userId,
  groupId
}) {
  const requestData = {
    groupId: groupId,
    userId: userId
  };
  console.log('Request data:', requestData);
  try {
    const response = await fetch('http://localhost:8080/api/groupUsers', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestData)
    });
    console.log('Response:', response);
    if (response.ok) {
      // Handle success response
      console.log('Group user added successfully!');
    } else {
      // Handle error response
      console.error('Failed to add group user:', response.statusText);
    }
  } catch (error) {
    console.error('Error adding group user:', error);
  }
}