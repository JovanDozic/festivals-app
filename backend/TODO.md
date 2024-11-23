can i get a list of employees that work on any festival that i'm organizer of?

--

[] get all employees for my organization: GET organizer/employee

- this will include festival IDs that they're working on so organizer can change it

[] get employees on the festival: GET organizer/festival/{festivalId}/employee

[] create employee: POST organizer/employee

[] get employee: GET organizer/employee/{employeeId}

[] update employee: PUT organizer/employee/{employeeId}

[] delete employee: DELETE organizer/{employeeId}

[] add employee to the festival: PUT organizer/festival/{festivalId}/employee/{employeeId}/employ

[] remove employee from the festival: PUT organizer/festival/{festivalId}/employee/{employeeId}/fire
