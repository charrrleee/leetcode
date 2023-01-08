select w2.Id
from Weather w1,
     Weather w2
where TO_DAYS(w2.RecordDate) - TO_DAYS(w1.RecordDate) = 1
  and w2.Temperature > w1.Temperature;