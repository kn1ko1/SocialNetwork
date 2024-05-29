const {
  useState
} = React;
// import { GroupDetails } from "./GroupDetails";

export function EventForm({
  group,
  socket
}) {
  const [dateTime, setDateTime] = useState("");
  const [description, setDescription] = useState("");
  const [title, setTitle] = useState("");

  // Handler for form submission
  const submit = async e => {
    e.preventDefault(); // Prevent page reload

    // Create a combined date-time string
    const dateTimeString = new Date(dateTime).toISOString();
    const formData = new FormData();

    // Append form data
    formData.append("dateTime", dateTimeString);
    formData.append("description", description);
    formData.append("groupId", group.groupId);
    formData.append("title", title);
    console.log("Event Form data being sent to backend: ", formData);
    try {
      // Send user data to the server
      await fetch("http://localhost:8080/api/events", {
        method: "POST",
        credentials: "include",
        body: formData
      });
      let obj = {
        code: 6,
        body: JSON.stringify(formData)
      };
      socket.send(JSON.stringify(obj));

      // Reset form fields after successful submission
      setDateTime("");
      setDescription("");
      setTitle("");
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
  }, /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal"
  }, "Post Event Here"), /*#__PURE__*/React.createElement("form", {
    onSubmit: submit
  }, /*#__PURE__*/React.createElement("div", {
    className: "form-floating mb-3"
  }, /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "eventFormTitle",
    placeholder: "Type your event title here...",
    onChange: e => setTitle(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "form-floating mb-3"
  }, /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "eventFormDescription",
    placeholder: "Type your event Description here...",
    onChange: e => setDescription(e.target.value)
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