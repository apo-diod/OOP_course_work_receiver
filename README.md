# OOP_course_work_receiver

Has 2 API Endpoints:
1) /add_module -> adds new module and hosts it]
  syntax: 
  {
	  "module": "flask",
	  "settings": {
		  "port": "port"
	  }
  }
2) /link -> links receiver and transformer modules
  syntax:
  {
	  "first": "receiver_id",
	  "second": "transformer_id"
  }
