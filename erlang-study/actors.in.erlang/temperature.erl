% Understanding actor concurrency, Part1: Actors in Erlang, https://www.infoworld.com/article/2077999/understanding-actor-concurrency--part-1--actors-in-erlang.html
-module(temperature).
-export([temperatureConverter/0]).


% Value = 4.
% Value = 6 
%  "=" is not an assignment operator at all in Erlang, it is really a pattern-matching operator. 
temperatureConverter() ->
	receive
		% words starting with lowercase letters are atoms -- fixed symbols 
		% that always represent a constant, number, or string of the same name, 
		% even across machines.
		{toF, C} ->
			io:format("~p C is ~p F~n", [C, 32+C*9/5]),
			temperatureConverter();
		{toC, F} -> 
			io:format("~p F is ~p C~n", [F, (F-32)*5/9]),
			temperatureConverter();
		{stop} ->
			io:format("Stopping~n");
		Other -> 
			io:format("Unkown: ~p~n", [Other]),
			temperatureConverter()
	end.
