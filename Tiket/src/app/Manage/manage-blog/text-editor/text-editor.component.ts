import { Component, OnInit, Input, ViewChild, ElementRef } from '@angular/core';

@Component({
  selector: 'app-text-editor',
  templateUrl: './text-editor.component.html',
  styleUrls: ['./text-editor.component.scss']
})
export class TextEditorComponent implements OnInit {
  @Input() c1: string;
  @Input() c2: string;

  @ViewChild('title', {static: true})
  private title: ElementRef;
  @ViewChild('content', {static: true})
  private content: ElementRef;
  constructor() {

  }

  ngOnInit() {
    
  }
  log(){
    console.log(this.title.nativeElement.innerHTML)
    this.content.nativeElement.innerHTML="gfffggfghfghfhgfhfgh"
  }
  format(cmnd) {
    document.execCommand(cmnd, false);
  }

  getTitle() {
    return this.title.nativeElement.innerHTML;
  }

  getContent() {
    return this.content.nativeElement.innerHTML;
  }
  setTitle(t:string){
    console.log
    this.title.nativeElement.innerHTML = t
  }
  setContent(t:string){
    this.content.nativeElement.innerHTML=t
  }

}
