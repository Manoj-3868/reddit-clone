import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, retry, throwError } from 'rxjs';
import { Subreddit } from '../Models/subreddit';

@Injectable({
  providedIn: 'root'
})
export class SubredditService {
  httpHeader={
    headers:new HttpHeaders({
      'Content-Type':'application/json'
    }
    )
  }
  baseUrl:string="http://localhost:3000"
  constructor(private httpClient:HttpClient) { }

  getAllSubreddits(): Observable<Array<Subreddit>> {
    return this.httpClient.get<Array<Subreddit>>(this.baseUrl+'/subreddit').pipe(
      retry(1),
      catchError(this.httpError)
    );
  }

  createSubreddit(subredditModel: Subreddit): Observable<Subreddit> {
    return this.httpClient.post<Subreddit>(this.baseUrl+'/subreddit',JSON.stringify(subredditModel),this.httpHeader).pipe(
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
