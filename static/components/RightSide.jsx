import React, { useState, useEffect } from "react";
import { Notification } from "./Notifications";

export default function RightSide(props) {
  useEffect(() => {
    const getAllNotifications = async () => {
      if (props["openConnection"] !== undefined) {
        const connectionResponse = await props["openConnection"]();
        if (connectionResponse === "connection open") {
          props["fetchRequestData"]();
        }
      }
    };
    getAllNotifications();
  }, []);
  return (
    <div
      id="rightSide"
      style={{
        height: props.page === "profiles" && "100px",
        gridColumn: props.page === "profiles" && "1 / span 4",
        gridRow: props.page === "profiles" && "1",
        margin: props.page === "profiles" && "20px",
      }}
    >
      <Notification page={props.page} />
    </div>
  );
}
