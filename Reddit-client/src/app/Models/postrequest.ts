export class Postrequest {
    postname: string;
    url: string;
    description: string;
    username: string;
    subredditname: string; 
    constructor(postname:string,url:string,description: string,username: string,subredditname: string){
        this.postname=postname
        this.url=url
        this.description=description
        this.username=username
        this.subredditname =subredditname
}
}