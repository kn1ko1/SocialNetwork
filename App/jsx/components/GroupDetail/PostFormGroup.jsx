const { useState } = React
import { GroupDetails } from "../../GroupDetails.js";

export function PostFormGroup({ group, fetchFunc }) {
	const [body, setBody] = useState("");
	const [selectedFile, setSelectedFile] = useState(null);

	// Handler for form submission
	const submit = async (e) => {
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
				body: formData,
			});

			// Reset form fields after successful submission
			setBody("");
			setSelectedFile(null);
			document.getElementById("postFormBody").value = "";

			const pageContainer = document.querySelector(".page-container");
			ReactDOM.render(<GroupDetails group={group} />, pageContainer)


		} catch (error) {
			console.error("Error submitting post:", error);
		}

		fetchFunc(group)
	};

	// Handler for file selection
	const handleFileChange = (e) => {
		setSelectedFile(e.target.files[0]);
	};

	const handleSelectFile = () => {
		const fileInput = document.getElementById("fileInput");
		fileInput.click();
	};

	return (
		<div>
			<main className="postForm container" style={{ maxWidth: "400px" }}>
				<h2 className="h3 mb-3 fw-normal" style={{ textDecoration: 'underline' }}>New Post</h2>
				<form onSubmit={submit}>
					<div className="mb-3">
						<input
							type="text"
							className="form-control"
							id="postFormBody"
							placeholder="Type your post here..."
							onChange={(e) => setBody(e.target.value)}
						/>
					</div>
					<div>
						<button
							type="button"
							className="btn btn-primary"
							style={{ marginRight: "10px" }}
							onClick={handleSelectFile}
						>
							Select File
						</button>
						<span>{selectedFile ? selectedFile.name : "No file selected"}</span>
						<input
							type="file"
							id="fileInput"
							accept="image/*"
							style={{ display: "none" }}
							onChange={handleFileChange}
						/>
					</div>
					<br /> {/* Line break */}
					<button className="w-100 btn btn-lg btn-primary" type="submit">
						Submit
					</button>
				</form>
			</main>
		</div>
	)
}