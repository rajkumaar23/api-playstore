<?php
/**
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */

use Carbon\Carbon;
use Symfony\Component\DomCrawler\Crawler;


/**
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */
class API
{
    private $playstoreURL;
    private $crawler;
    private $conn;
    private $data;

    public function __construct($package)
    {
        $this->playstoreURL = 'https://play.google.com/store/apps/details?id=' . $package;
        $this->conn = (new DB())->getConn();
        $this->data = $this->conn->findOne(['packageID' => $package]);
        if (empty($this->data) || Utils::shouldUpdateCache($this->data['lastCached'])) {
            if ($html = file_get_contents($this->playstoreURL)) {
                $this->crawler = new Crawler($html);
                $htlgb = $this->crawler->filter('.htlgb');
                $this->data['packageID'] = $package;
                $this->data['version'] = $htlgb->eq(6)->text();
                $this->data['installs'] = $htlgb->eq(5)->text();
                $this->data['size'] = $htlgb->eq(3)->text();
                $this->data['lastUpdated'] = $htlgb->eq(1)->text();
                $this->data['rating'] = $this->crawler->filter('.BHMmbe')->eq(0)->text();
                $this->data['noOfUsersRated'] = filter_var($this->crawler->filter('.EymY4b')->eq(0)->text(), FILTER_SANITIZE_NUMBER_INT);
                $this->data['developer'] = $htlgb->eq(sizeof($htlgb) == 20 ? 17 : 18)->text();
                $this->data['lastCached'] = Carbon::now()->toISOString();
                $this->conn->insertOne($this->data);
            } else {
                throw new Exception("Invalid Package ID");
            }
        }
    }

    /**
     * @return mixed
     */
    public function getPackageID()
    {
        return $this->data['packageID'];
    }

    /**
     * @return string
     */
    public function getVersion()
    {
        return $this->data['version'];
    }

    /**
     * @return string
     */
    public function getInstalls()
    {
        return $this->data['installs'];
    }

    /**
     * @return string
     */
    public function getSize()
    {
        return $this->data['size'];
    }

    /**
     * @return string
     */
    public function getLastUpdated()
    {
        return $this->data['lastUpdated'];
    }

    /**
     * @return string
     */
    public function getRating()
    {
        return $this->data['rating'];
    }

    /**
     * @return string
     */
    public function getNoOfUsersRated()
    {
        return $this->data['noOfUsersRated'];
    }

    /**
     * @return string
     */
    public function getDeveloper()
    {
        return $this->data['developer'];
    }
}
