import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'clockclient';
  clockHtml: SafeHtml | undefined;

  constructor(
    private http:HttpClient,
    private sanitizer:DomSanitizer
  ){}

  ngOnInit(){
    this.http.get('/api/clockWidget',{responseType:'text'}).subscribe(res=>{
      this.clockHtml = this.sanitizer.bypassSecurityTrustHtml(res);
    })
  }
}
