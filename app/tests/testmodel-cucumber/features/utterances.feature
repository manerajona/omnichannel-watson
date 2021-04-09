Feature: Infoguia utterances
  In order to determine if the assistant is behaving correctly
  We need to check the following utterances

  Scenario:
    When user asks to watson: "¿Dónde puedo pedir una pizza?"
    Then watson should respond with intent: "ver_gastronomia"
  
  Scenario:
    When user asks to watson: "que bancos hay abiertos"
    Then watson should respond with intent: "ver_bancos"

  Scenario:
    When user asks to watson: "¿Existe Dios?"
    Then watson should respond with intent: ""

  Scenario:
    When user asks to watson: "¿Qué iglesias hay?"
    Then watson should respond with intent: "ver_templos"

  Scenario:
    When user asks to watson: "Ik wil een biertje!"
    Then watson should respond with intent: ""

  Scenario:
    When user asks to watson: "¿Qué lugar vende cerveza?"
    Then watson should respond with intent: "ver_gastronomia"

  Scenario:
    When user asks to watson: "Menú"
    Then watson should respond with intent: "ver_menu"

  Scenario:
    When user asks to watson: "Mostrar contactos de emergencia"
    Then watson should respond with intent: "ver_emergencias"

  Scenario:
    When user asks to watson: "ver menú"
    Then watson should respond with intent: "ver_menu"

  Scenario:
    When user asks to watson: "guía de comercios"
    Then watson should respond with intent: "ver_comercios"

  Scenario:
    When user asks to watson: "cosas de hogar"
    Then watson should respond with intent: "ver_hogar"

  Scenario:
    When user asks to watson: "que entidades hay?"
    Then watson should respond with intent: "ver_entidades"

  Scenario:
    When user asks to watson: "mostrame bancos"
    Then watson should respond with intent: "ver_bancos"

  Scenario:
    When user asks to watson: "para hacer turismo?"
    Then watson should respond with intent: "ver_lugares_turisticos"

  Scenario:
    When user asks to watson: "la iglesia?"
    Then watson should respond with intent: "ver_templos"

  Scenario:
    When user asks to watson: "donde atienden emergencias???"
    Then watson should respond with intent: "ver_emergencias"

  Scenario:
    When user asks to watson: "sos"
    Then watson should respond with intent: "ver_emergencias"

  Scenario:
    When user asks to watson: "donde encuentro farmacias abiertas"
    Then watson should respond with intent: "ver_farmacias"

  Scenario:
    When user asks to watson: "farmacia de turno"
    Then watson should respond with intent: "ver_farmacias"

  Scenario:
    When user asks to watson: "quiero ofertas!"
    Then watson should respond with intent: "ver_ofertas"

  Scenario:
    When user asks to watson: "necesito descuentos!!"
    Then watson should respond with intent: "ver_ofertas"

  Scenario:
    When user asks to watson: "algun evento"
    Then watson should respond with intent: "ver_eventos"

  Scenario:
    When user asks to watson: "¿Qué eventos hay el finde?"
    Then watson should respond with intent: "ver_eventos"

  Scenario:
    When user asks to watson: "números de remis"
    Then watson should respond with intent: "ver_remis"

  Scenario:
    When user asks to watson: "taxis disponibles?"
    Then watson should respond with intent: "ver_remis"
