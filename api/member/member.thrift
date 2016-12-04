namespace go member

struct Member {
  1: string _id,
  2: string line_id
  3: string name,
  4: string pic,
  5: string message,
  6: string day
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
  list<Member> getNotAssign(),
  void addMember(1:Member member),
}
