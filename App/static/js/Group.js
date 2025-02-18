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

    const groupData = {
      creatorId: currentUserId,
      description: description,
      title: title
    };
    let obj = {
      code: 10,
      body: JSON.stringify(groupData)
    };
    socket.send(JSON.stringify(obj));
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
  const groupStyle = {
    maxWidth: '1300px',
    background: 'linear-gradient(to bottom, #c7ddef, #ffffff)',
    // Light blue/grey to white gradient
    borderRadius: '10px',
    boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
    // Optional: Add shadow for depth
    padding: '40px',
    margin: 'auto',
    marginBottom: '20px',
    // Adjust spacing between post cards
    border: '1px solid #ccc' // Add a thin border
  };
  const opaqueStyle = {
    backgroundColor: 'rgba(255, 255, 255, 0.25)',
    // Adjust the opacity here 
    maxWidth: '1300px',
    borderRadius: '10px',
    boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
    // Optional: Add shadow for depth
    padding: '40px',
    margin: 'auto',
    marginBottom: '20px' // Adjust spacing between post cards
  };
  const overriddenStyle = {
    ...groupStyle,
    maxWidth: '400px' // Override maxWidth
  };
  return /*#__PURE__*/React.createElement("div", {
    className: "container",
    style: opaqueStyle
  }, selectedGroup ? /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("button", {
    onClick: () => setSelectedGroup(null),
    type: "submit",
    className: "btn btn-primary",
    style: {
      margin: '30px'
    }
  }, "Go Back"), /*#__PURE__*/React.createElement(GroupDetails, {
    group: selectedGroup,
    socket: socket
  })) : /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("form", {
    onSubmit: create,
    className: "container",
    style: overriddenStyle
  }, /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("h2", {
    style: {
      textDecoration: 'underline',
      textAlign: 'center'
    }
  }, "New Group"), /*#__PURE__*/React.createElement("label", {
    htmlFor: "exampleTitle",
    className: "form-label"
  }, "Title"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "exampleTitle",
    "aria-describedby": "emailHelp",
    value: title,
    placeholder: "Type your group title here...",
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
    placeholder: "Type your group description here...",
    onChange: e => setDescription(e.target.value)
  })), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary"
  }, "Create")), /*#__PURE__*/React.createElement("div", {
    className: "text-center",
    style: opaqueStyle
  }, /*#__PURE__*/React.createElement("h2", {
    style: {
      textDecoration: 'underline'
    }
  }, "Groups"), /*#__PURE__*/React.createElement("div", {
    id: "groupData",
    className: "row"
  }, groupData !== null ? groupData.map(group => /*#__PURE__*/React.createElement("div", {
    key: group.title,
    className: "col-lg-3 col-md-4 col-sm-6 col-12 mb-4",
    onClick: () => handleGroupClick(group)
  }, /*#__PURE__*/React.createElement("div", {
    style: groupStyle
  }, /*#__PURE__*/React.createElement("h3", {
    style: {
      textDecoration: 'underline',
      fontSize: '22px'
    }
  }, group.title), /*#__PURE__*/React.createElement("p", {
    style: {
      fontSize: '18px'
    }
  }, group.description)))) : /*#__PURE__*/React.createElement("div", {
    id: "noGroupsError"
  }, "There are no created groups yet")))));
}