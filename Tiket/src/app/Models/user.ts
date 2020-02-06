export class User {
    private firstname:string
    private lastname:string
    private email:string
    private phone:string
    private password:string

    constructor(fname:string, lname:string,mail:string,phone:string,pass:string){
        this.firstname =fname;
        this.lastname =lname;
        this.email =mail;
        this.phone =phone;
        this.password =pass;
    }
    public getFirstname(){
        return this.firstname
    }
    public getLastname(){
        return this.lastname
    }
    public getEmail(){
        return this.email
    }
    public getPhone(){
        return this.phone
    }
    public getPassword(){
        return this.password
    }
}
