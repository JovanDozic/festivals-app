can i get a list of employees that work on any festival that i'm organizer of?

--

[] create employee: POST organizer/employee

- ovo ce da se kreira novi nalog i profil
- onda ti treba i endpoint call da ga zaposlis

[] get all employees for my organization: GET organizer/employee

- this will include festival IDs that they're working on so organizer can change it

[] get employees on the festival: GET organizer/festival/{festivalId}/employee

[] get employee: GET organizer/employee/{employeeId}

[] update employee: PUT organizer/employee/{employeeId}

[] delete employee: DELETE organizer/{employeeId}

[] add employee to the festival: PUT organizer/festival/{festivalId}/employee/{employeeId}/employ

[] remove employee from the festival: PUT organizer/festival/{festivalId}/employee/{employeeId}/fire