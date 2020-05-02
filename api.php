<?php

/**
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */

use Symfony\Component\DomCrawler\Crawler;

class API
{
    private $conn;
    private $data;

    public function __construct($package)
    {
        $this->conn = (new DB())->getConn();
        $this->data = $this->conn->findOne(['packageID' => $package]);
        if (!empty($this->data)) {
            if (Utils::shouldUpdateCache($this->data['lastCached'])) {
                $protocol = (!empty($_SERVER['HTTPS']) && $_SERVER['HTTPS'] !== 'off' || $_SERVER['SERVER_PORT'] == 443) ? "https://" : "http://";
                $command = "curl -s " . $protocol . $_SERVER['HTTP_HOST'] . "/update-cache?id=$package";
                exec("nohup '.$command.' > /dev/null 2>&1 & echo $!");
            }
        } else {
            $this->data = self::updateCache($package);
        }
    }

    public static function updateCache($package)
    {
        $playstoreURL = 'https://play.google.com/store/apps/details?id=' . $package;
        if ($html = file_get_contents($playstoreURL)) {
            $conn = (new DB())->getConn();
            $crawler = new Crawler($html);
            $htlgb = $crawler->filter('.htlgb');
            $data['packageID'] = $package;
            $data['version'] = $htlgb->eq(6)->text();
            $data['installs'] = $htlgb->eq(5)->text();
            $data['size'] = $htlgb->eq(3)->text();
            $data['lastUpdated'] = $htlgb->eq(1)->text();
            $data['rating'] = $crawler->filter('.BHMmbe')->eq(0)->text();
            $data['noOfUsersRated'] = filter_var($crawler->filter('.EymY4b')->eq(0)->text(), FILTER_SANITIZE_NUMBER_INT);
            $data['developer'] = $htlgb->eq(sizeof($htlgb) == 20 ? 17 : 18)->text();
            $data['lastCached'] = time();
            $conn->updateOne(['packageID' => $package], ['$set' => $data], ['upsert' => true]);
            return $data;
        } else {
            throw new Exception("Invalid Package ID");
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
