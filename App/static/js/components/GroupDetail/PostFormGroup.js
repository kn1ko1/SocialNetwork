const {
  useState
} = React;
import { GroupDetails } from "../../GroupDetails.js";
export function PostFormGroup({
  group,
  fetchFunc
}) {
  const [body, setBody] = useState("");
  const [selectedFile, setSelectedFile] = useState(null);

  // Handler for form submission
  const submit = async e => {
    e.preventDefault(); // Prevent page reload

    const formData = new FormData();
    // Handle body content
    let requestBody = body;
    if (selectedFile && !requestBody.trim()) {
      // If selectedFile is present and body is empty or whitespace only,
      // set requestBody to a space character
      requestBody = " ";
    }
    // Append form data
    formData.append("body", requestBody);
    formData.append("groupId", group.groupId);
    if (selectedFile) {
      formData.append("image", selectedFile);
    }
    console.log("Form data being sent to backend: ", formData);
    try {
      // Send user data to the server
      await fetch("http://localhost:8080/api/posts", {
        method: "POST",
        credentials: "include",
        body: formData
      });

      // Reset form fields after successful submission
      setBody("");
      setSelectedFile(null);
      document.getElementById("postFormBody").value = "";
      const pageContainer = document.querySelector(".page-container");
      ReactDOM.render( /*#__PURE__*/React.createElement(GroupDetails, {
        group: group
      }), pageContainer);
    } catch (error) {
      console.error("Error submitting post:", error);
    }
    fetchFunc(group);
  };

  // Handler for file selection
  const handleFileChange = e => {
    setSelectedFile(e.target.files[0]);
  };
  const handleSelectFile = () => {
    const fileInput = document.getElementById("fileInput");
    fileInput.click();
  };
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("main", {
    className: "postForm container",
    style: {
      maxWidth: "400px"
    }
  }, /*#__PURE__*/React.createElement("h2", {
    className: "h3 mb-3 fw-normal",
    style: {
      textDecoration: 'underline'
    }
  }, "New Post"), /*#__PURE__*/React.createElement("form", {
    onSubmit: submit
  }, /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "postFormBody",
    placeholder: "Type your post here...",
    onChange: e => setBody(e.target.value)
  })), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("button", {
    type: "button",
    className: "btn btn-primary",
    style: {
      marginRight: "10px"
    },
    onClick: handleSelectFile
  }, "Select File"), /*#__PURE__*/React.createElement("span", null, selectedFile ? selectedFile.name : "No file selected"), /*#__PURE__*/React.createElement("input", {
    type: "file",
    id: "fileInput",
    accept: "image/*",
    style: {
      display: "none"
    },
    onChange: handleFileChange
  })), /*#__PURE__*/React.createElement("br", null), " ", /*#__PURE__*/React.createElement("button", {
    className: "w-100 btn btn-lg btn-primary",
    type: "submit"
  }, "Submit"))));
}