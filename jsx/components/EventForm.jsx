const { useState } = React
// import { GroupDetails } from "./GroupDetails";

export function EventForm({ group }) {
    const [dateTime, setDateTime] = useState("");
    const [description, setDescription] = useState("");
    const [title, setTitle] = useState("");
    // Handler for form submission
    const submit = async (e) => {
        e.preventDefault(); // Prevent page reload
 
 
        const formData = new FormData();
 
 
        // Append form data
        formData.append("dateTime", "2022-01-02");
        formData.append("description", description);
        formData.append("groupId", group.groupId);
        formData.append("title", title);
 
 
        console.log("Form data being sent to backend: ", formData);
 
 
        try {
            // Send user data to the server
            await fetch("http://localhost:8080/api/events", {
                method: "POST",
                credentials: "include",
                body: formData,
            });
 
 
            // Reset form fields after successful submission
            setDateTime("");
            setDescription("");
            setTitle("");
            document.getElementById("eventFormDescription").value = "";
            document.getElementById("eventFormTitle").value = "";
        } catch (error) {
            console.error("Error submitting event:", error);
        }
        // const pageContainer = document.querySelector(".page-container")
        // ReactDOM.render(<GroupDetails group={group} />, pageContainer)
    };
 
 
    return (
        <div>
            <main className="eventForm container" style={{ maxWidth: "400px" }}>
                <h1 className="h3 mb-3 fw-normal">Post Event Here</h1>
                <form onSubmit={submit}>
                    <div className="form-floating mb-3">
                        <input
                            type="text"
                            className="form-control"
                            id="eventFormTitle"
                            placeholder="Type your event title here..."
                            onChange={(e) => setTitle(e.target.value)}
                        />
                    </div>
                    <div className="form-floating mb-3">
                        <input
                            type="text"
                            className="form-control"
                            id="eventFormDescription"
                            placeholder="Type your event Description here..."
                            onChange={(e) => setDescription(e.target.value)}
                        />
                    </div>
                    <br />
                    <button className="w-100 btn btn-lg btn-primary" type="submit">
                        Submit
                    </button>
                </form>
            </main>
        </div>
    )
 }
 