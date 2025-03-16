import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, retry, throwError } from 'rxjs';
import { Vote } from '../Models/vote';

@Injectable({
  providedIn: 'root'
})
export class VoteService {
  httpHeader={
    headers:new HttpHeaders({
      'Content-Type':'application/json'
    }
    )
  }
  baseUrl:string="http://localhost:3000"


  constructor(private httpClient:HttpClient) { }

  vote(votePayload: Vote): Observable<any> {
    console.log('vote-',votePayload)
    return this.httpClient.post(this.baseUrl+'/vote',JSON.stringify(votePayload),this.httpHeader).pipe(
      retry(1),
      catchError(this.httpError)
    );
  }


  httpError(error:HttpErrorResponse){
    let msg='';
    if(error.error instanceof ErrorEvent){
      msg=error.error.message;
    }
    else{
      msg=`Error Code:${error.status}\nMessafe:${error.message}`;
    }
    console.log(msg);
    return throwError(msg);
  }
}
