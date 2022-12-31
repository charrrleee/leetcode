select p1.Email
from Person p1
         Join Person p2 on p1.Id <> p2.Id
where p1.Email = p2.Email
group by p1.Email;