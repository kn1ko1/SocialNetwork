const {
  useState,
  useEffect
} = React;
import { GroupDetails } from "./GroupDetails.js";
import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js";
export const renderGroup = ({
  socket
}) => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Group, {
    socket: socket
  }), pageContainer);
};
export function Group({
  socket
}) {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [groupData, setGroupData] = useState([]);
  const [selectedGroup, setSelectedGroup] = useState(null);
  //const [showGroupDetails, setShowGroupDetails] = useState(false);
  const {
    currentUserId
  } = getCurrentUserId();
  const fetchGroupData = async () => {
    try {
      const promises = [];
      promises.push(fetch("http://localhost:8080/api/groups"));
      promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/groupUsers`));
      const results = await Promise.all(promises);
      const groupListResponse = results[0];
      const joinedGroupsResponse = results[1];
      if (!groupListResponse.ok) {
        throw new Error('Failed to fetch group list');
      }
      if (!joinedGroupsResponse.ok) {
        throw new Error('Failed to fetch joined group list');
      }
      const groupListData = await groupListResponse.json();
      const joinedGroupsData = await joinedGroupsResponse.json();
      if (groupListData != null) {
        for (let i = 0; i < groupListData.length; i++) {
          if (joinedGroupsData !== null) {
            const joinedGroup = joinedGroupsData.find(group => group.groupId === groupListData[i].groupId);
            // If a corresponding group is found in joinedGroupsData
            if (joinedGroup) {
              // Add a new field 'isMember' to groupListData and set its value to true
              groupListData[i].isMember = true;
            } else {
              // If no corresponding group is found, set 'isMember' to false or undefined
              groupListData[i].isMember = false;
            }
          } else {
            groupListData[i].isMember = false;
          }
        }
      }
      setGroupData(groupListData);
    } catch (error) {
      console.error('Error fetching group data:', error);
    }
  };
  useEffect(() => {
    if (currentUserId) {
      fetchGroupData();
    }
  }, [currentUserId]);
  const create = async e => {
    e.preventDefault(); // prevent reload.

    const groupData = new FormData();

    // Append form data
    groupData.append("group-title", title);
    groupData.append("group-description", description);
    console.log("Group data being sent to backend:", title);
    console.log("Group data being sent to backend:", description);

    // Send user data to golang api/PostHandler.go.
    await fetch("http://localhost:8080/api/groups", {
      method: "POST",
      credentials: "include",
      body: groupData
    });
    setTitle("");
    setDescription("");
    document.getElementById("exampleTitle").value = "";
    document.getElementById("exampleDescription").value = "";
    fetchGroupData();
  };
  const handleGroupClick = group => {
    setSelectedGroup(group);
    //setShowGroupDetails(true);
  };
  const handleGoBack = () => {
    setSelectedGroup(null);
    setShowGroupDetails(false); // Update showGroupDetails to false when going back
  };
  return /*#__PURE__*/React.createElement("div", null, selectedGroup ? /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("button", {
    onClick: () => setSelectedGroup(null)
  }, "Go Back"), /*#__PURE__*/React.createElement(GroupDetails, {
    group: selectedGroup,
    socket: socket
  })) : /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("form", {
    onSubmit: create,
    className: "container",
    style: {
      maxWidth: "400px"
    }
  }, /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "exampleTitle",
    className: "form-label"
  }, "Title"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "exampleTitle",
    "aria-describedby": "emailHelp",
    value: title,
    onChange: e => setTitle(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "exampleInputPassword1",
    className: "form-label"
  }, "Description"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "exampleDescription",
    value: description,
    onChange: e => setDescription(e.target.value)
  })), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary"
  }, "Create")), /*#__PURE__*/React.createElement("div", {
    id: "groupData"
  }, groupData !== null ? groupData.map(group => /*#__PURE__*/React.createElement("div", {
    key: group.title,
    onClick: () => handleGroupClick(group)
  }, /*#__PURE__*/React.createElement("h3", null, group.title), /*#__PURE__*/React.createElement("p", null, group.description))) : /*#__PURE__*/React.createElement("div", {
    id: "noGroupsError"
  }, "There are no created groups yet"))));
}