import { Component, OnInit } from '@angular/core';
import * as L from 'leaflet'
@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.scss']
})
export class MapComponent implements OnInit {
  private map: any

  constructor() { }

  ngOnInit() {
    this.map = L.map('map').setView([42.35, -71.08], 13);

    // load a tile layer
    L.tileLayer('http://tiles.mapc.org/basemap/{z}/{x}/{y}.png',
      {
        attribution: 'Tiles by <a href="http://mapc.org">MAPC</a>, Data by <a href="http://mass.gov/mgis">MassGIS</a>',
        maxZoom: 17,
        minZoom: 9
      }).addTo(this.map);
  }

}
