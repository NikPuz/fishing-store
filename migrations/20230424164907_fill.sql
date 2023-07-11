-- +goose Up

--
-- таблица `categories`
--

INSERT INTO public.categories (id, name) VALUES (0, 'Без категории');

--
-- таблица `manufacturers`
--

INSERT INTO public.manufacturers (id, name) VALUES (0, 'Без производителя');