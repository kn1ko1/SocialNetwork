import { formattedDate } from "../shared/FormattedDate.js";
export function GroupDetailsEvents({
  groupEvents
}) {
  return /*#__PURE__*/React.createElement("div", {
    className: "groupEvents"
  }, /*#__PURE__*/React.createElement("h2", null, "Events"), groupEvents !== null && groupEvents.length > 0 ? groupEvents.map((event, index) => /*#__PURE__*/React.createElement("div", {
    key: index,
    className: "row"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col-3"
  }, /*#__PURE__*/React.createElement("p", {
    className: "d-inline-flex gap-1"
  }, /*#__PURE__*/React.createElement("button", {
    className: "btn btn-primary",
    type: "button",
    "data-bs-toggle": "collapse"
    // Use unique ID for data-bs-target
    ,
    "data-bs-target": `#collapseExample${index}`,
    "aria-expanded": "false",
    "aria-controls": `collapseExample${index}`
  }, event.title))), /*#__PURE__*/React.createElement("div", {
    className: "col-9"
  }, /*#__PURE__*/React.createElement("div", {
    className: "collapse",
    id: `collapseExample${index}`
  }, /*#__PURE__*/React.createElement("div", {
    className: "card card-body"
  }, event.description, " - ", formattedDate(event.dateTime)))))) : /*#__PURE__*/React.createElement("p", null, "No Events"));
}