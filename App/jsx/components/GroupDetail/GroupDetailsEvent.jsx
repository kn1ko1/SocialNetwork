import { formattedDate } from "../shared/FormattedDate.js"

export function GroupDetailsEvents({groupEvents}) {
   
    
    return (
        <div className="groupEvents">
          <h2 style={{ textDecoration: 'underline' }}>Events</h2>
          {groupEvents !== null && groupEvents.length > 0 ? (
            <div className="accordion" id="eventAccordion">
              {groupEvents.map((event, index) => (
                <div key={index} className="accordion-item">
                  <h2 className="accordion-header" id={`heading${index}`}>
                    <button
                      className="accordion-button collapsed"
                      type="button"
                      data-bs-toggle="collapse"
                      data-bs-target={`#collapse${index}`}
                      aria-expanded="false"
                      aria-controls={`collapse${index}`}
                    >
                      {event.title}
                    </button>
                  </h2>
                  <div
                    id={`collapse${index}`}
                    className="accordion-collapse collapse"
                    aria-labelledby={`heading${index}`}
                    data-bs-parent="#eventAccordion"
                  >
                    <div className="accordion-body">
                      <p>{event.description}</p>
                      <small>{formattedDate(event.dateTime)}</small>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <p>No Events</p>
          )}
        </div>
      );
}