const { useState } = React;
import { initializeSocket } from "./app.js";
import { renderNavbar } from "./components/shared/Navbar.js";
import { renderHome } from "./Home.js";
import { renderLogin } from "./Login.js";

export const renderRegister = () => {
	const pageContainer = document.querySelector(".page-container");
	ReactDOM.render(<Register />, pageContainer);
};

export function Register() {
	const [formValues, setFormValues] = useState({
		email: "",
		password: "",
		firstName: "",
		lastName: "",
		dob: "",
		username: "",
		bio: "",
		isPublic: true,
	});
	const [selectedFile, setSelectedFile] = useState(null);
	const [isRegistered, setIsRegistered] = useState(false);
	const [errors, setErrors] = useState({});

	const validate = () => {
		const errors = {};
		const emailPattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
		const dobDate = new Date(formValues.dob);
		const currentDate = new Date();
		if (!emailPattern.test(formValues.email)) {
			errors.email = "Please enter a valid email address following the format: email@example.com";
		}
		if (!formValues.password) {
			errors.password = "Password is required.";
		}
		if (!formValues.firstName) {
			errors.firstName = "First name is required.";
		}
		if (!formValues.lastName) {
			errors.lastName = "Last name is required.";
		}
		if (isNaN(dobDate.getTime()) || dobDate >= currentDate) {
			errors.dob = "Please enter a valid Date of Birth in the past.";
		}
		if (formValues.username.length < 4 || formValues.username.length > 50) {
			errors.username = "Username must be between 4 and 50 characters.";
		}
		return errors;
	};

	const handleChange = (e) => {
		const { name, value, type, checked } = e.target;
		setFormValues((prevValues) => ({
			...prevValues,
			[name]: type === "checkbox" ? checked : type === "radio" ? JSON.parse(value) : value,
		}));
	};

	const handleFileChange = (e) => {
		setSelectedFile(e.target.files[0]);
	};

	const handleSelectFile = () => {
		const fileInput = document.getElementById("fileInput");
		fileInput.click();
	};

	const submit = async (e) => {
		e.preventDefault(); // prevent reload.

		const validationErrors = validate();
		if (Object.keys(validationErrors).length > 0) {
			setErrors(validationErrors);
			return;
		}

		try {
			const formData = new FormData();
			Object.keys(formValues).forEach((key) => {
				formData.append(key, formValues[key]);
			});
			if (selectedFile) {
				formData.append("image", selectedFile);
			}
			console.log("formData:", formData);
			// Send user data to backend
			const response = await fetch("http://localhost:8080/auth/registration", {
				method: "POST",
				body: formData,
			});
			console.log("Response:", response)
			if (!response.ok) {
				setErrors({ server: data.message || "Registration failed" });
				throw new Error("Invalid credentials");
			}

			const data = await response.json();
			console.log(data)
			if (data.success) { 
				setIsRegistered(true) 	
			} else {
				const validationErrors = {}
				if (data.errorField == "Email") {
					validationErrors.email = data.errorMessage
				}
				if (data.errorField == "Username") {
					validationErrors.username = data.errorMessage
				}
				if (Object.keys(validationErrors).length > 0) {
					setErrors(validationErrors);
					return;
				}
				return;
			}
		} catch (error) {
			console.error("Registration error:", error);
		}
	};

	if (isRegistered) {
		const socket = initializeSocket();
		renderNavbar({ socket });
		renderHome({ socket });
	}

	return (
		<div className="container login-container" style={{ minHeight: "100vh", display: "flex", justifyContent: "center", alignItems: "center", paddingBottom: "20px" }}>
			<div className="logo-container">
				<img src="../static/sphere-logo.png" alt="Logo" className="logo" />
			</div>

			<h1 className="h3 mb-3 fw-normal login-text">Register</h1>
			<form onSubmit={submit}>
				<div className="row mb-3">
					<div className="col-md-6">
						<label htmlFor="floatingInput" className="form-label">Email address</label>
						<input
							required
							type="email"
							className={`form-control ${errors.email ? "is-invalid" : ""}`}
							id="floatingInput"
							name="email"
							placeholder="name@example.com"
							onChange={handleChange}
						/>
						{errors.email && <div className="invalid-feedback">{errors.email}</div>}
					</div>
					<div className="col-md-6">
						<label htmlFor="regpassword" className="form-label">Password</label>
						<input
							required
							type="password"
							className={`form-control ${errors.password ? "is-invalid" : ""}`}
							id="regpassword"
							name="password"
							placeholder="Password"
							onChange={handleChange}
						/>
						{errors.password && <div className="invalid-feedback">{errors.password}</div>}
					</div>
				</div>

				<div className="row mb-3">
					<div className="col-md-6">
						<label htmlFor="firstName" className="form-label">First Name</label>
						<input
							required
							type="text"
							className={`form-control ${errors.firstName ? "is-invalid" : ""}`}
							id="firstName"
							name="firstName"
							placeholder="John"
							onChange={handleChange}
						/>
						{errors.firstName && <div className="invalid-feedback">{errors.firstName}</div>}
					</div>
					<div className="col-md-6">
						<label htmlFor="lastName" className="form-label">Last Name</label>
						<input
							required
							type="text"
							className={`form-control ${errors.lastName ? "is-invalid" : ""}`}
							id="lastName"
							name="lastName"
							placeholder="Doe"
							onChange={handleChange}
						/>
						{errors.lastName && <div className="invalid-feedback">{errors.lastName}</div>}
					</div>
				</div>

				<div className="row mb-3">
					<div className="col-md-6">
						<label htmlFor="dob" className="form-label">Date of Birth</label>
						<input
							required
							type="date"
							className={`form-control ${errors.dob ? "is-invalid" : ""}`}
							id="dob"
							name="dob"
							placeholder="16/01/1998"
							onChange={handleChange}
						/>
						{errors.dob && <div className="invalid-feedback">{errors.dob}</div>}
					</div>
					<div className="col-md-6">
						<label htmlFor="username" className="form-label">Username</label>
						<input
							required
							type="text"
							className={`form-control ${errors.username ? "is-invalid" : ""}`}
							id="username"
							name="username"
							placeholder="Johnny"
							onChange={handleChange}
						/>
						{errors.username && <div className="invalid-feedback">{errors.username}</div>}
					</div>
				</div>

				<div className="row mb-3">
					<div className="col-md-6">
						<label htmlFor="image" className="form-label">Avatar Image (optional)</label>
						<div className="input-group">
							<input
								type="file"
								id="fileInput"
								accept="image/*"
								style={{ display: "none" }}
								onChange={handleFileChange}
							/>
							<button
								className="btn btn-primary rounded"
								type="button"
								onClick={handleSelectFile}
							>
								Select File
							</button>
							<span className="input-group-text" style={{ backgroundColor: "transparent", border: "none" }}>{selectedFile ? selectedFile.name : "No file selected"}</span>
						</div>
					</div>







					<div className="col-md-6">
						<div className="mb-3">
							<label className="form-label">Profile Visibility</label>
							<div className="form-check">
								<input
									className="form-check-input"
									type="radio"
									id="public-status"
									name="isPublic"
									value={true}
									checked={formValues.isPublic === true}
									onChange={handleChange}
								/>
								<label className="form-check-label" htmlFor="public-status">
									Public
								</label>
							</div>
							<div className="form-check">
								<input
									className="form-check-input"
									type="radio"
									id="private-status"
									name="isPublic"
									value={false}
									checked={formValues.isPublic === false}
									onChange={handleChange}
								/>
								<label className="form-check-label" htmlFor="private-status">
									Private
								</label>
							</div>
						</div>
					</div>
				</div>

				<div className="row mb-3">
					<div className="col-md-12">
						<label htmlFor="bio" className="form-label">About me (optional)</label>
						<input
							type="text"
							className="form-control"
							id="bio"
							name="bio"
							placeholder="About Me"
							onChange={handleChange}
						/>
					</div>
				</div>

				<button className="btn btn-primary" type="submit">Register</button>

			</form>

			<div className="error-message"></div>
			<br />

			<div className="mb-3">
				<span className="login-text">Already have an account? &nbsp;</span>
				<button type="button" className="btn btn-primary" onClick={renderLogin}>Log in</button>
			</div>
		</div>
	);
}