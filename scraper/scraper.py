import datetime
import logging

import requests
import datetime
import url
import bs4


class Vacancy:
    name = str
    city = str
    description = str
    company_name = str
    publication_date_text = str
    scrape_date = datetime.date

    def __init__(self,
                 name,
                 city,
                 description,
                 company_name,
                 publication_date_text,
                 scrape_date):
        self.name = name
        self.city = city
        self.description = description
        self.company_name = company_name
        self.publication_date_text = publication_date_text
        self.scrape_date = scrape_date


class DocumentElement:
    element_attributes = dict

    def __init__(self, element_attributes: dict):
        self.element_attributes = element_attributes


class VacanciesItemMapping:
    name_field = DocumentElement
    city_field = DocumentElement
    description_field = DocumentElement
    company_name_field = DocumentElement
    publication_date_field = DocumentElement

    def __init__(self,
                 name_field: DocumentElement,
                 city_field: DocumentElement,
                 description_field: DocumentElement,
                 company_name_field: DocumentElement,
                 publication_date_field: DocumentElement):
        self.name_field = name_field
        self.city_field = city_field
        self.description_field = description_field
        self.company_name_field = company_name_field
        self.publication_date_field = publication_date_field


class VacanciesSearchResultPageMapping:
    vacancy_block_element = str
    vacancy_block_class = str

    def __init__(self, vacancy_block_element, vacancy_block_class):
        self.vacancy_block_element = vacancy_block_element
        self.vacancy_block_class = vacancy_block_class


class VacanciesSearchResultPage:
    page_data = str

    def __init__(self, page_data):
        self.page_data = page_data


class VacanciesScrapingParams:
    query = dict
    headers = dict
    url = str

    def __init__(self, query, headers, url):
        self.query = query
        self.headers = headers
        self.url = url


class VacanciesScraper:
    session = requests.Session
    logger = logging.Logger

    def __init__(self, session, logger):
        self.session = session
        self.logger = logger

    def parse(self, params: VacanciesScrapingParams, headers: dict):
        built_url = url.add_url_params(url=params.url, params=params.query)
        response = self.session.get(built_url, headers=headers)
        if response.status_code != 200:
            self.logger.log(msg=f"Error getting the page {built_url}", level=logging.ERROR)
            self.logger.log(msg=f"{response.text}", level=logging.ERROR)
        return VacanciesSearchResultPage(response.text)

    def extract_vacancies(self,
                          page: VacanciesSearchResultPage,
                          page_mapping: VacanciesSearchResultPageMapping,
                          item_mapping: VacanciesItemMapping):
        page_data = page.page_data
        soup = bs4.BeautifulSoup(page_data, 'html.parser')

        results = []
        scrape_date = datetime.date.today()


        for element in soup.find_all(page_mapping.vacancy_block_element, class_=page_mapping.vacancy_block_class):
            name = ""
            city = ""
            description = ""
            company_name = ""
            publication_date_text = ""

            name_element = element.find(attrs=item_mapping.name_field.element_attributes)
            if name_element:
                name = name_element.text.strip()

            city_element = element.find(attrs=item_mapping.city_field.element_attributes)
            if city_element:
                city = city_element.text.strip()

            description_element = element.find(attrs=item_mapping.description_field.element_attributes)
            if description_element:
                description = description_element.text.strip()

            company_name_element = element.find(attrs=item_mapping.company_name_field.element_attributes)
            if company_name_element:
                company_name = company_name_element.text.strip()

            publication_date_text_element = element.find(attrs=item_mapping.publication_date_field.element_attributes)
            if publication_date_text_element:
                publication_date_text = publication_date_text_element.text.strip()

            vacancy = Vacancy(name, city, description, company_name, publication_date_text, scrape_date)
            results.append(vacancy)

        return results