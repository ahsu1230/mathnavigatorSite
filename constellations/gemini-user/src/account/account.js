"use strict";
require("./account.sass");
import React from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import API from "../utils/api.js";

export class AccountPage extends React.Component {
    state = {
        id: 1,
        
        primaryEmail: "",
        users: [],
        transactions: [],
        userClasses: [],
    }
    
    componentDidMount = () => {
        const id = this.state.id;
        
        if (id) {
            API.get("api/accounts/account/" + id)
                .then((res) => this.fetchData(res))
                .catch((err) => this.fetchDataError(err));
        }
    }
    
    fetchData = (res) => {
        const id = res.data.id;
        this.setState({
            id: id,
            primaryEmail: res.data.primaryEmail,
            users: [],
            transactions: []
        });
        
        Promise.all([
            API.get("api/classes/all"),
            API.get("api/programs/all"),
            API.get("api/semesters/all"),
        ]).then((res) => {
            const allClasses = res[0].data;
            const allPrograms = res[1].data;
            const allSemesters = res[2].data;
            
            console.log(allPrograms)
            
            API.get("api/users/account/" + id)
                .then((res) => {
                    const users = res.data
                    
                    this.setState({ users: users })
                    
                    let userClasses = [];
                    users.map((user, index) => {
                        API.get("api/user-classes/user/" + user.id).then((res) => {
                            if (res.data.length) {
                                let classes = res.data.map((c, index) => {
                                    let matchedClass = allClasses.find(element => element.classId == c.classId)
                                    let matchedProgram = allPrograms.find(element => element.programId == matchedClass.programId)
                                    let matchedSemester = allSemesters.find(element => element.semesterId == matchedClass.semesterId)
                                    return matchedProgram.name + ' ' + matchedSemester.title
                                })
                                console.log(classes)
                                
                                userClasses.push({
                                    id: user.id,
                                    name: user.firstName + ' ' + user.lastName,
                                    classes: classes
                                });
                                this.setState({ userClasses: userClasses })
                            }
                        })
                    })
                })
        });
            
        API.get("api/transactions/account/" + id)
            .then((res) => this.setState({ transactions: res.data }))            
    };
    
    formatCurrency = (amount) => {
        return new Intl.NumberFormat("en-US", {
            style: "currency",
            currency: "USD",
        }).format(amount);
    };
    
    renderClassList = (classes) => {
        return classes.map((c, index) => {
            return <p key={index} className="classList-item">{c}</p>
        })
    }
    
    render = () => {
        console.log(this.state)
        
        const usersList = this.state.users.map((user, index) => {
            return (
                <ul key={index} className="has-edit">
                    <li>{user.firstName + ' ' + user.lastName}</li>
                    <li>{user.email}</li>
                    <Link to="" className="edit">Edit</Link>
                </ul>
        )})
        
        let balance = 0;
        const transactionsList = this.state.transactions.map((transaction, index) => {
            balance += parseInt(transaction.amount)
            return (
                <ul key={index}>
                    <li>{transaction.paymentType}</li>
                    <li>{this.formatCurrency(transaction.amount)}</li>
                    <li>{transaction.paymentNotes}</li>
                </ul>
        )})
        
        const classRegistrationList = this.state.userClasses.map((user, index) => {
            return (
                <div key={index}>
                    <p>{user.name}</p>
                    {this.renderClassList(user.classes)}
                </div>
        )})
        
        return (
            <div id="view-account">
                <h1>Your Account</h1>
                
                <div>
                    <p>
                        <span>Primary email: {this.state.primaryEmail}</span>
                        <Link to="" className="edit">Edit</Link>
                    </p>
                    <p>
                        <Link to="">Change password...</Link>
                    </p>
                </div>
                
                <div>
                    <h2>User Information</h2>
                    {usersList}
                </div>
                
                <div>
                    <h2>Account Balance & History</h2>
                    {transactionsList}
                    <hr />
                    <span>Account balance: {this.formatCurrency(balance)}</span>
                </div>
                
                <div>
                    <h2>Class Registrations</h2>
                    {classRegistrationList}
                </div>
                
                <div>
                    <Link to="" className="red">Request to Delete Account</Link>
                </div>
                
            </div>
        );
    }
}
