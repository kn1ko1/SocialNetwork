export const formattedDate = (milliseconds) => {
    const date = new Date(milliseconds);
    const options = {
        year: "numeric",
        month: "numeric",
        day: "numeric",
        hour: "numeric",
        minute: "numeric",
        hour12: true // Use 12-hour format
    };
    return date.toLocaleString(undefined, options);
};
