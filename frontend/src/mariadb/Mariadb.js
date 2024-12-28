import React, { useState } from 'react';
import './Mariadb.css';

function MariaDB() {
  const [users, setUsers] = useState([]);
  const [user, setUser] = useState({ name: '', age: '' });
  const [error, setError] = useState(null);
  const [showTable, setShowTable] = useState(false);

  const fetchUsers = async () => {
    try {
      const response = await fetch('http://localhost:8080/users');
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      const data = await response.json();
      console.log('Fetched users:', data); // Debugging log
      setUsers(data); // Update state with fetched data
      setShowTable(true); // Show table if users are fetched
    } catch (error) {
      console.error('Error fetching users:', error);
      setError('Error fetching users');
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setUser({ ...user, [name]: value });
  };

  const addUser = async () => {
    try {
      const response = await fetch('http://localhost:8080/insert', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name: user.name, age: parseInt(user.age, 10) }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      const data = await response.json();
      setUsers([...users, data]); // Add new user to the list
      setUser({ name: '', age: '' }); // Reset user input
      await fetchUsers(); // Fetch users again to update the list
    } catch (error) {
      console.error('Error adding user:', error);
      setError('Error adding user');
    }
  };
  

  const deleteUser = async (name, age) => {
    try {
      const response = await fetch('http://localhost:8080/delete', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, age }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      // Update the users state to reflect the deletion
      setUsers(users.filter((user) => user.name !== name || user.age !== age));
      await response.json();
      await fetchUsers(); // Fetch users again to update the list
    } catch (error) {
      console.error('Error deleting user:', error);
      setError('Error deleting user');
    }
  };

  return (
    <div className='container'>
      <div className='centered-div'><h2>MariaDB</h2>
      <p>MariaDB is a community-developed fork of MySQL.</p>
      <center>
        <input
          type="text"
          name="name"
          value={user.name}
          onChange={handleInputChange}
          placeholder="Name"
        />
        <input
          type="number"
          name="age"
          value={user.age}
          onChange={handleInputChange}
          placeholder="Age"
        />
        <button onClick={addUser} >Add User</button>    
        </center>
        <center>
      <button onClick={fetchUsers}>Get Users</button></center>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {showTable && users.length > 0 && (
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th>Age</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user) => (
              <tr key={user.ID}>
                <td>{user.Name}</td>
                <td>{user.Age}</td>
                <td>
                  <button onClick={() => deleteUser(user.Name, user.Age)}>Delete</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
      {showTable && users.length === 0 && (
        <p>No users found.</p>
      )}
    </div>
    </div>
  );
}

export default MariaDB;
