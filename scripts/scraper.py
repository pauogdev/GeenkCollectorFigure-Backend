# coding: utf-8
# Script base para scraping de tiendas de figuras
# Puedes implementar aquí la lógica para obtener datos de las tiendas

# Ejemplo de estructura:
# - Obtener HTML
# - Parsear datos
# - Guardar en MongoDB

import requests
from bs4 import BeautifulSoup
from pymongo import MongoClient
import os
import json

# Configuración de MongoDB
MONGO_URI = os.getenv('MONGO_URI', 'mongodb://localhost:27017')
DB_NAME = os.getenv('DB_NAME', 'GeenkCollectorFigure')
COLLECTION_NAME = 'Figura'

client = MongoClient(MONGO_URI)
db = client[DB_NAME]
figures_collection = db[COLLECTION_NAME]

def scrap_amiami():
    url = 'https://www.amiami.com/eng/search/list/?s_originaltitle_id=0&s_sortkey=price_asc&s_st_condition_flg=1&s_cate3=0&s_cate2=0&s_cate1=0&s_condition_flg=1&s_page=1'
    response = requests.get(url)
    soup = BeautifulSoup(response.text, 'html.parser')
    items = soup.select('.itemlist .item')
    if not items:
        print('No se encontraron figuras con el selector .itemlist .item. Revisa el selector CSS.')
        return
    for item in items:
        nombre = item.select_one('.item_name').text if item.select_one('.item_name') else ''
        precio = item.select_one('.item_price').text if item.select_one('.item_price') else ''
        imagen = item.select_one('img')['src'] if item.select_one('img') else ''
        url_figura = item.select_one('a')['href'] if item.select_one('a') else ''
        marca = '' # Si la web lo muestra
        serie = '' # Si la web lo muestra
        tamano = '' # Si la web lo muestra
        figura = {
            'nombre': nombre,
            'precio': precio,
            'imagen': imagen,
            'url': url_figura,
            'marca': marca,
            'serie': serie,
            'tamano': tamano
        }
        result = figures_collection.insert_one(figura)
        print(f'Insertada figura: {nombre} (ID: {result.inserted_id})')

def scrap_amiami_api(gcode_list):
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36'
    }
    for gcode in gcode_list:
        url = 'https://api.amiami.com/api/v1.0/item'
        payload = {'gcode': gcode, 'lang': 'eng'}
        response = requests.get(url, params=payload, headers=headers)
        if response.status_code != 200:
            print(f'Error al consultar {gcode}: {response.status_code}')
            continue
        data = response.json()
        if 'item' not in data:
            print(f'No se encontró información para {gcode}')
            continue
        item = data['item']
        figura = {
            'nombre': item.get('name', ''),
            'precio': item.get('price', ''),
            'imagen': item.get('main_image', ''),
            'url': item.get('url', ''),
            'marca': item.get('maker_name', ''),
            'serie': item.get('original_title', ''),
            'tamano': item.get('size', '')
        }
        result = figures_collection.insert_one(figura)
        print(f'Insertada figura: {figura["nombre"]} (ID: {result.inserted_id})')

def get_gcodes_from_listing(pages=1):
    gcodes = set()
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36'
    }
    for page in range(1, pages+1):
        url = f'https://www.amiami.com/eng/search/list/?s_cate3=0&s_cate2=0&s_cate1=0&s_condition_flg=1&s_page={page}'
        response = requests.get(url, headers=headers)
        soup = BeautifulSoup(response.text, 'html.parser')
        # Buscar enlaces a la ficha de producto que contienen el gcode
        for a in soup.find_all('a', href=True):
            href = a['href']
            if '/eng/detail/?gcode=FIGURE-' in href:
                gcode = href.split('gcode=')[-1].split('&')[0]
                gcodes.add(gcode)
    return list(gcodes)

if __name__ == '__main__':
    try:
        from amiami import AmiAmi
    except ImportError:
        print('Instala la librería amiami: pip install amiami')
        exit(1)

    client_api = AmiAmi(lang='eng')
    # Puedes ajustar los parámetros de búsqueda según la documentación de la librería
    results = client_api.search(category='figure', page=1)
    print(f'Encontradas {len(results)} figuras en la primera página.')
    for item in results:
        figura = {
            'nombre': item.get('name', ''),
            'precio': item.get('price', ''),
            'imagen': item.get('main_image', ''),
            'url': item.get('url', ''),
            'marca': item.get('maker_name', ''),
            'serie': item.get('original_title', ''),
            'tamano': item.get('size', '')
        }
        result = figures_collection.insert_one(figura)
        print(f'Insertada figura: {figura["nombre"]} (ID: {result.inserted_id})')
    print('Scraping finalizado.')
