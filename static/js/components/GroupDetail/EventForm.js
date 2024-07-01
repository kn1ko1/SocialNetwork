import { getCurrentUserId } from "../shared/GetCurrentUserId.js";
const {
  useState
} = React;
// import { GroupDetails } from "./GroupDetails";

//Test Description
export function EventForm({
  group,
  socket
}) {
  const [dateTime, setDateTime] = useState("");
  const [eventDescription, setEventDescription] = useState("");
  const [eventTitle, setEventTitle] = useState("");
  const {
    currentUserId
  } = getCurrentUserId();

  // Handler for form submission
  const submit = async e => {
    e.preventDefault(); // Prevent page reload

    // Create a combined date-time string
    const dateTimeMillis = new Date(dateTime).getTime();
    const eventData = {
      dateTime: dateTimeMillis,
      description: eventDescription,
      groupId: group.groupId,
      title: eventTitle,
      userId: currentUserId
    };
    console.log("Event Form data being sent to backend: ", eventData);
    try {
      let obj = {
        code: 6,
        body: JSON.stringify(eventData)
      };
      socket.send(JSON.stringify(obj));

      // Reset form fields after successful submission
      setDateTime("");
      setEventDescription("");
      setEventTitle("");
      document.getElementById("eventFormDescription").value = "";
      document.getElementById("eventFormTitle").value = "";
    } catch (error) {
      console.error("Error submitting event:", error);
    }
  };
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("main", {
    className: "eventForm container",
    style: {
      maxWidth: "400px"
    }
  }, /*#__PURE__*/React.createElement("h2", {
    className: "h3 mb-3 fw-normal",
    style: {
      textDecoration: 'underline'
    }
  }, "New Event"), /*#__PURE__*/React.createElement("form", {
    onSubmit: submit
  }, /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "exampleTitle",
    className: "form-label"
  }, "Event Title"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "eventFormTitle",
    placeholder: "Title here...",
    onChange: e => setEventTitle(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "exampleDescription",
    className: "form-label"
  }, "Event Description"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "eventFormDescription",
    placeholder: "Description here...",
    onChange: e => setEventDescription(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "dateTime"
  }, "Date and Time of Event"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "datetime-local",
    className: "form-control reginput",
    id: "dateTime",
    onChange: e => setDateTime(e.target.value)
  })), /*#__PURE__*/React.createElement("br", null), /*#__PURE__*/React.createElement("button", {
    className: "w-100 btn btn-lg btn-primary",
    type: "submit"
  }, "Submit"))));
}