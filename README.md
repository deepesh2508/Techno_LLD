Backend: Go
DB: PostgreSQL
Cache / state: Redis
Auth: JWT (access + refresh)
AI: OpenAI (evaluation only)



Database Design
1. users
users (
  id UUID PK,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  createdby int
  updatedby int
  email TEXT UNIQUE,
  google_id TEXT UNIQUE NULL,
  name TEXT,
  avatar_url TEXT
)


2. user_auth_sessions (
  id UUID PK,
  created_at TIMESTAMP,
  user_id UUID FK(users.id),
  refresh_token TEXT,
  expires_at TIMESTAMP,
)

3. email_otps (
  id UUID PK,
  email TEXT,
  otp_hash TEXT,
  expires_at TIMESTAMP,
  used BOOLEAN DEFAULT false
)

4. questions (
  id UUID PK,
  created_at TIMESTAMP,
  title TEXT,
  description TEXT,
  difficulty ENUM('easy','medium','hard'),
  type ENUM('LLD','HLD'),
  acceptance_rate INT,
  is_locked BOOLEAN,
  is_coming_soon BOOLEAN,
)

5. question_companies (
  id UUID PK,
  question_id UUID FK(questions.id),
  company_name TEXT,
  company_logo_url TEXT
)

6. clarification_rules (
  id UUID PK,
  question_id UUID,
  intent TEXT,
  response TEXT,
  max_questions INT
)

7. interview_sessions (
  id UUID PK,
  user_id UUID,
  question_id UUID,
  status ENUM('clarifying','design','submitted'),
  clarifications_used INT,
  started_at TIMESTAMP,
  ended_at TIMESTAMP
)

8. clarification_logs (
  id UUID PK,
  session_id UUID,
  user_question TEXT,
  interviewer_response TEXT,
  created_at TIMESTAMP
)

9. solutions (
  id UUID PK,
  session_id UUID,
  functional_req TEXT,
  non_functional_req TEXT,
  entities TEXT,
  classes_and_relationships TEXT,
  API's TEXT,
  code TEXT,
  extra TEXT,
  language ENUM('go','cpp','java','python'),
  submitted_at TIMESTAMP
)

10. evaluations (
  id UUID PK,
  solution_id UUID,
  score INT,
  strengths TEXT,
  weaknesses TEXT,
  improvement_points TEXT,
  evaluated_at TIMESTAMP
)


API Design-->
1. Request email OTP
2. Verify OTP
3. Get profile
4. Update profile
5. Login
6. Logout
7. Get all questions
8. Get question
9. Start question
10. Clarify requirements
11. Proceed to design
12. Submit Solution
13. Get evaluation result


Submission
   ↓
Job Queue
   ↓
GPT Evaluation
   ↓
Store result
