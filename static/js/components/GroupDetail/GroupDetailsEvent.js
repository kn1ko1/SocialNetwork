import { formattedDate } from "../shared/FormattedDate.js";
export function GroupDetailsEvents({
  groupEvents
}) {
  return /*#__PURE__*/React.createElement("div", {
    className: "groupEvents"
  }, /*#__PURE__*/React.createElement("h2", {
    style: {
      textDecoration: 'underline'
    }
  }, "Events"), groupEvents !== null && groupEvents.length > 0 ? /*#__PURE__*/React.createElement("div", {
    className: "accordion",
    id: "eventAccordion"
  }, groupEvents.map((event, index) => /*#__PURE__*/React.createElement("div", {
    key: index,
    className: "accordion-item"
  }, /*#__PURE__*/React.createElement("h2", {
    className: "accordion-header",
    id: `heading${index}`
  }, /*#__PURE__*/React.createElement("button", {
    className: "accordion-button collapsed",
    type: "button",
    "data-bs-toggle": "collapse",
    "data-bs-target": `#collapse${index}`,
    "aria-expanded": "false",
    "aria-controls": `collapse${index}`
  }, event.title)), /*#__PURE__*/React.createElement("div", {
    id: `collapse${index}`,
    className: "accordion-collapse collapse",
    "aria-labelledby": `heading${index}`,
    "data-bs-parent": "#eventAccordion"
  }, /*#__PURE__*/React.createElement("div", {
    className: "accordion-body"
  }, /*#__PURE__*/React.createElement("p", null, event.description), /*#__PURE__*/React.createElement("small", null, formattedDate(event.dateTime))))))) : /*#__PURE__*/React.createElement("p", null, "No Events"));
}