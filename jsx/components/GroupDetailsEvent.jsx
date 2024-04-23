export function GroupDetailsEvents({groupEvents}) {
    return (
        <div className="groupEvents">
            <h2>Events</h2>
            {groupEvents !== null && groupEvents.length > 0 ? (
                groupEvents.map((event, index) => (
                    <div key={index} className="row">
                        <div className="col-3">
                            <p className="d-inline-flex gap-1">
                                <button
                                    className="btn btn-primary"
                                    type="button"
                                    data-bs-toggle="collapse"
                                    // Use unique ID for data-bs-target
                                    data-bs-target={`#collapseExample${index}`}
                                    aria-expanded="false"
                                    aria-controls={`collapseExample${index}`}
                                >
                                    {event.title}
                                </button>
                            </p>
                        </div>
                        <div className="col-9">
                            <div className="collapse" id={`collapseExample${index}`}>
                                <div className="card card-body">
                                    {event.description} - {event.dateTime}
                                </div>
                            </div>
                        </div>
                    </div>
                ))
            ) : (
                <p>No Events</p>
            )}
        </div>)
}