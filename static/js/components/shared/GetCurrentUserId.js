const {
  useState,
  useEffect
} = React;
export const getCurrentUserId = () => {
  const [currentUserId, setCurrentUserId] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  useEffect(() => {
    const fetchUserId = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/userId", {
          credentials: "include"
        });
        if (response.ok) {
          const userId = await response.json();
          setCurrentUserId(userId);
        } else {
          setError("Failed to fetch userId");
        }
      } catch (error) {
        setError("Error fetching userId");
      } finally {
        setIsLoading(false);
      }
    };
    fetchUserId();
  }, []);
  return {
    currentUserId,
    isLoading,
    error
  };
};