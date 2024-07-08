import { getCurrentUserId } from "../shared/GetCurrentUserId.js";

const { useState } = React
// import { GroupDetails } from "./GroupDetails";

//Test Description
export function EventForm({ group, socket, fetchFunc }) {
    const [dateTime, setDateTime] = useState("");
    const [eventDescription, setEventDescription] = useState("");
    const [eventTitle, setEventTitle] = useState("");
    const [attendance, setAttendance] = useState("attending");
    const { currentUserId } = getCurrentUserId();


    // Handler for form submission
    const submit = async (e) => {
        e.preventDefault(); // Prevent page reload

        // Create a combined date-time string
        const dateTimeMillis = new Date(dateTime).getTime();

        const eventData = {
            dateTime: dateTimeMillis,
            description: eventDescription,
            groupId: group.groupId,
            title: eventTitle,
            userId: currentUserId,
            attendance: attendance // Include attendance in event data
        };

        console.log("Event Form data being sent to backend: ", eventData);

        try {

            let obj = { code: 6, body: JSON.stringify(eventData) }
            socket.send(JSON.stringify(obj));

            // Reset form fields after successful submission
            setDateTime("");
            setEventDescription("");
            setEventTitle("");
            setAttendance("attending"); // Reset to default attendance
            document.getElementById("eventFormDescription").value = "";
            document.getElementById("eventFormTitle").value = "";
            document.getElementById("attending").checked = true; // Reset radio button
        } catch (error) {
            console.error("Error submitting event:", error);
        }
        fetchFunc(group)
    };

    return (
        <div>
            <main className="eventForm container" style={{ maxWidth: "400px" }}>
                <h2 className="h3 mb-3 fw-normal" style={{textDecoration: 'underline'}}>New Event</h2>
                <form onSubmit={submit}>
                
                    <div className="mb-3">
                    <label htmlFor="exampleTitle" className="form-label">
					Event Title
				  </label>
                        <input
                            type="text"
                            className="form-control"
                            id="eventFormTitle"
                            placeholder="Title here..."
                            onChange={(e) => setEventTitle(e.target.value)}
                        />
                    </div>

                    <div className="mb-3">

                    <label htmlFor="exampleDescription" className="form-label">
					Event Description
				  </label>
                        <input
                            type="text"
                            className="form-control"
                            id="eventFormDescription"
                            placeholder="Description here..."
                            onChange={(e) => setEventDescription(e.target.value)}
                        />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="dateTime">Date and Time of Event</label>
                        <input
                            required
                            type="datetime-local"
                            className="form-control reginput"
                            id="dateTime"
                            onChange={(e) => setDateTime(e.target.value)}
                        />
                    </div>

                    <div className="mb-3">
            <div>
              <div className="form-check">
                <input
                  className="form-check-input"
                  type="radio"
                  name="attendanceOptions"
                  id="attending"
                  value="attending"
                  checked={attendance === 'attending'}
                  onChange={(e) => setAttendance(e.target.value)}
                />
                <label className="form-check-label" htmlFor="attending">
                  Attending
                </label>
              </div>
              <div className="form-check">
                <input
                  className="form-check-input"
                  type="radio"
                  name="attendanceOptions"
                  id="notAttending"
                  value="notAttending"
                  checked={attendance === 'notAttending'}
                  onChange={(e) => setAttendance(e.target.value)}
                />
                <label className="form-check-label" htmlFor="notAttending">
                  Not Attending
                </label>
              </div>
            </div>
          </div>


                    <br />
                    <button className="w-100 btn btn-lg btn-primary" type="submit">
                        Submit
                    </button>
                </form>
            </main>
        </div>
    );
}
