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
    this.http.get('api/clockWidget/clockHtml',{responseType:'text'}).subscribe(res=>{
      this.clockHtml = this.sanitizer.bypassSecurityTrustHtml(res);
      this.loadScript('api/clockWidget/clockJs');
    });
  }

  public loadScript(url: string) {
    const body = <HTMLDivElement> document.body;
    const script = document.createElement('script');
    script.innerHTML = '';
    script.src = url;
    script.async = false;
    body.appendChild(script);
  }
}
