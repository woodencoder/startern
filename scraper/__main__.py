import sys

from scraper import *


def main(args):
    session = requests.session()
    scraper = VacanciesScraper(session, logging.getLogger())
    scraping_params = VacanciesScrapingParams(
        {"area": "1", "enable_snippets": "true", "st": "searchVacancy", "text": "стажер"},
        {},
        "https://hh.ru/search/vacancy"
    )
    search_results_page = scraper.parse(
        scraping_params,
        headers={
            "cache-control": "max-age=0",
            "sec-ch-ua": "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"",
            "sec-ch-ua-mobile": "?0",
            "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36",
            "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
            "sec-fetch-site": "none",
            "sec-fetch-mode": "navigate",
            "sec-fetch-user": "?1",
            "sec-fetch-dest": "document",
            "accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7"
        })

    page_mapping = VacanciesSearchResultPageMapping("div", "vacancy-serp-item")
    item_mapping = VacanciesItemMapping(
        name_field=DocumentElement({"data-qa": "vacancy-serp__vacancy-title"}),
        city_field=DocumentElement({"data-qa": "vacancy-serp-item__meta-info"}),
        description_field=DocumentElement({"data-qa": "vacancy-serp__vacancy_snippet_responsibility"}),
        company_name_field=DocumentElement({"data-qa": "vacancy-serp__vacancy-employer"}),
        publication_date_field=DocumentElement(
            {"class": "vacancy-serp-item__publication-date vacancy-serp-item__publication-date_long"})
    )

    vacancies = scraper.extract_vacancies(search_results_page, page_mapping, item_mapping)

    for element in vacancies:
        print("====")
        print(f"{element.company_name}")
        print(f"{element.name}")
        print(f"{element.description}")
        print("====")


main(args=sys.argv)