import { Component, inject, OnInit } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { FestivalService } from '../../services/festival/festival.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
  standalone: true,
  imports: [MatCardModule],
})
export class HomeComponent implements OnInit {
  private festivalService = inject(FestivalService);
  private router = inject(Router);

  festivalCount = 0;
  attendeesCount = 0;

  ngOnInit(): void {
    this.festivalService.getFestivalsCount().subscribe((response) => {
      this.festivalCount = response;
    });

    this.festivalService.getAttendeesCount().subscribe((response) => {
      this.attendeesCount = response;
    });
  }

  async onClick() {
    if (!(await this.router.navigate(['login']))) {
      if (!(await this.router.navigate(['festivals']))) {
        if (!(await this.router.navigate(['organizer/my-festivals']))) {
          if (!(await this.router.navigate(['employee/my-festivals']))) {
            this.router.navigate(['admin/users']);
          }
        }
      }
    }
  }
}
