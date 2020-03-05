import { ApolloService } from 'src/app/Services/apollo.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-blog-detail',
  templateUrl: './blog-detail.component.html',
  styleUrls: ['./blog-detail.component.scss']
})
export class BlogDetailComponent implements OnInit {
  private blogs:any;
  private recs:any;
  constructor(private activatedRoute: ActivatedRoute,
              private router: Router,
              private apollo:ApolloService) { 

    this.activatedRoute.queryParams.subscribe(async a=>{
      await this.getParam(a)
    })
  }

  ngOnInit() {
  }
  getParam(a:any){
    this.apollo.searchBlogById(parseInt(a.id)).subscribe(async b => {
      await this.getData(b)
    })
    this.apollo.searchRecBlog(parseInt(a.id)).subscribe(async b =>{
      await this.getRecs(b)
    })
  }

  getData(b:any){
    this.blogs = b.data.getblogbyid
  }
  getRecs(b:any){
    
    this.recs = b.data.getrecblog
    console.log(this.recs)
  }

  fb() {
    window.open('http://www.facebook.com/sharer.php?u=localhost:4200/Blog/Detail?id=' + this.blogs.id, 'facebookShare', 'width=626,height=436');
  }

  wa() {
    window.open('https://api.whatsapp.com/send?text=localhost:4200/Blog/Detail?id=' + this.blogs.id)
  }

  copy() {
    navigator.clipboard.writeText('localhost:4200/Blog/Detail?id=' + this.blogs.id);
  }

  line() {
    window.open('http://line.me/R/msg/text/?Blog/Detail?id=' + this.blogs.id);
  }
  ToDetail(b: any) {
    this.router.navigate(['/Blog/Detail'], {
      queryParams: {
        id: b.id,
      }
    });
  }

}
