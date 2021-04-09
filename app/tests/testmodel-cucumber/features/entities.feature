Feature: Infoguia Entities
  In order to determine if the assistant is behaving correctly
  We need to check the following utterances

  Scenario:
    When user asks to watson: "1"
    Then watson should respond with entity: "1"

  Scenario:
    When user asks to watson: "99"
    Then watson should respond with entity: "99"

  Scenario:
    When user asks to watson: "uno"
    Then watson should respond with entity: "1"

  Scenario:
    When user asks to watson: "dos"
    Then watson should respond with entity: "2"

  Scenario:
    When user asks to watson: "tres"
    Then watson should respond with entity: "3"

  Scenario:
    When user asks to watson: "cuatro"
    Then watson should respond with entity: "4"

  Scenario:
    When user asks to watson: "cinco"
    Then watson should respond with entity: "5"

  Scenario:
    When user asks to watson: "seis"
    Then watson should respond with entity: "6"

  Scenario:
    When user asks to watson: "siete"
    Then watson should respond with entity: "7"

  Scenario:
    When user asks to watson: "ocho"
    Then watson should respond with entity: "8"

  Scenario:
      When user asks to watson: "SI"
      Then watson should respond with entity: "si"

  Scenario:
      When user asks to watson: "NO"
      Then watson should respond with entity: "no"