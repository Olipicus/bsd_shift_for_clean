namespace go member

struct Member {
  1: string _id,
  2: string name,
  3: string pic,
  4: string message,
  5: string day
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
