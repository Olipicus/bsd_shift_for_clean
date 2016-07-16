namespace go member

const string MongoAddress = "127.0.0.1:27017"
const string DatabaseName = "bsd_shift_for_clean"
const string Collection = "member"

struct Member {
  1: string _id,
  2: string name,
  3: string pic,
  4: string day
}

struct ResultDay {
  1: string day,
  2: string color,
  3: list<Member> members
}

service MemberService {
  list<Member> assignDay(1:string id),
  list<ResultDay> getResults(),
  Member getMember(1:string id),
  ResultDay getResultByDay(1:string day),
}
