const { useState, useEffect } = React

export function PostFormGroup({ group }) {
	const [body, setBody] = useState("");
	const [selectedFile, setSelectedFile] = useState(null);

	// Handler for form submission
	const submit = async (e) => {
		e.preventDefault(); // Prevent page reload

		const formData = new FormData();

		// Append form data
		formData.append("body", body);
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
				<h1 className="h3 mb-3 fw-normal">Post Message Here</h1>
				<form onSubmit={submit}>
					<div className="form-floating mb-3">
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