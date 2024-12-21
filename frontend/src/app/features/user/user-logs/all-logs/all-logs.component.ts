import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatMenuModule } from '@angular/material/menu';
import { MatTableModule } from '@angular/material/table';
import { Log } from '../../../../models/common/log.model';
import { LogsService } from '../../../../services/user/logs.service';
import { AuthService } from '../../../../services/auth/auth.service';
import { FormsModule } from '@angular/forms';
import { ViewLogComponent } from '../view-log/view-log.component';

@Component({
  selector: 'app-all-logs',
  imports: [
    CommonModule,
    FormsModule,
    MatButtonModule,
    MatTooltipModule,
    MatIconModule,
    MatDialogModule,
    MatDividerModule,
    MatCardModule,
    MatChipsModule,
    MatMenuModule,
    MatTableModule,
  ],
  templateUrl: './all-logs.component.html',
  styleUrls: ['./all-logs.component.scss', '../../../../app.component.scss'],
})
export class AllLogsComponent implements OnInit {
  private logService = inject(LogsService);
  private authService = inject(AuthService);
  private dialog = inject(MatDialog);

  filterOptions: string[] = ['All'];
  selectedChip = 'All';

  userRole = '';
  logs: Log[] = [];
  displayedColumns = ['id', 'username', 'createdAt', 'message', 'actions'];

  ngOnInit() {
    this.userRole = this.authService.getUserRole() ?? 'ADMINISTRATOR';
    this.setFilterOptions();
    this.logService.getLogs().subscribe((logs) => (this.logs = logs));
  }

  private setFilterOptions() {
    const roleFilters = {
      ADMINISTRATOR: [
        'All',
        'System',
        'Administrators',
        'Attendees',
        'Employees',
        'Organizers',
      ],
      EMPLOYEE: ['All', 'Employees', 'Attendees'],
      ORGANIZER: ['All', 'Organizers', 'Employees', 'Attendees'],
    };
    this.filterOptions =
      roleFilters[this.userRole as keyof typeof roleFilters] ??
      this.filterOptions;
  }

  get filteredLogs(): Log[] {
    if (this.selectedChip === 'All') {
      return this.logs;
    }
    if (this.selectedChip === 'System') {
      return this.logs.filter((log) => !log.role);
    }
    return this.logs.filter(
      (log) => log.role === this.selectedChip.slice(0, -1).toUpperCase(),
    );
  }

  onViewClick(log: Log) {
    this.dialog.open(ViewLogComponent, {
      data: { log: log },
      width: '400px',
      height: 'auto',
      disableClose: true,
    });
  }
}
