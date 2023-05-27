create table personalities
(
  id serial primary key,
  name varchar,
  history varchar
);

INSERT INTO personalities
  (name, history)
VALUES
  ('Maria da Silva', 'Conhecida como "Tia Maria", foi uma figura carismática e querida no bairro. Ela dedicou sua vida ao trabalho comunitário, ajudando crianças carentes por meio de projetos educacionais e culturais. Sua bondade e generosidade deixaram uma marca duradoura na comunidade.'),
  ('João Santos', 'Mais conhecido como "João do Mercado", ele foi um empreendedor local que construiu um pequeno mercado na região. Com seu carisma e preços justos, o estabelecimento se tornou um ponto de encontro para os moradores, além de fornecer empregos para muitos jovens do bairro.'),
  ('Ana Costa', 'Conhecida como "Dona Ana", foi uma ativista ambiental que lutou pela preservação das áreas verdes de Campo Grande. Ela liderou campanhas de conscientização sobre a importância do meio ambiente e organizou mutirões de limpeza para manter as praças e parques do bairro limpos e saudáveis. Seu trabalho incansável inspirou muitos a se preocuparem com a natureza local.');